package main

import (
	"log"
	"net"
	"net/rpc"

	nm "github.com/efournival/ter-lri/go-numeric-monoid"
)

// MaxTasksRPC will set the maximum number of tasks that can be exchanged through RPC's stealing requests
const MaxTasksRPC = 200

type (
	// Server is the RPC server type that is able to handle steal and sync requests
	Server struct {
		RPC        *rpc.Server
		Address    string
		TaskStream chan nm.GoMonoid
		Sync       chan chan nm.MonoidResults
	}

	// StealReply are the arguments to answer a steal request
	StealReply struct {
		Count int
		Tasks [MaxTasksRPC]nm.MonoidStorage
	}
)

// NewServer returns a new server automatically connected to the specified address
func NewServer(address string, ts chan nm.GoMonoid, sc chan chan nm.MonoidResults) *Server {
	s := &Server{rpc.NewServer(), address, ts, sc}

	if err := s.RPC.RegisterName("Danse", s); err != nil {
		log.Panicln("RPC register failed:", err)
	}

	// TODO: ensure this is working for multiple clients
	go func() {
		conn, err := net.Listen("tcp", s.Address)

		if err != nil {
			log.Panicln("Listener accept failed:", err)
		}

		s.RPC.Accept(conn)
	}()

	return s
}

// StealRequest is the RPC handler to a steal request
func (s *Server) StealRequest(max int, reply *StealReply) error {
	var tasks []nm.MonoidStorage

	for i := 0; i < max; i++ {
		select {
		// We can steal a task, add it to our steal answer
		case t := <-s.TaskStream:
			tasks = append(tasks, t.GetBytes())
			t.Free()
		// No task left
		// TODO: add timer channel
		default:
			break
		}
	}

	reply.Count = len(tasks)
	copy(reply.Tasks[:], tasks)

	return nil
}

// SyncRequest is the RPC handler to a sync request
func (s *Server) SyncRequest(_ bool, reply *nm.MonoidResults) error {
	ch := make(chan nm.MonoidResults, 1)
	s.Sync <- ch
	*reply = <-ch

	if debug {
		log.Println("Syncing", *reply)
	}

	return nil
}
