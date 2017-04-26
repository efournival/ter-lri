package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

const (
	MAX_TASKS_RPC = 100
	STEAL_COUNT   = 100
)

type (
	Server struct {
		RPC        *rpc.Server
		Address    string
		TaskStream chan nm.GoMonoid
		Sync       chan chan nm.MonoidResults
	}

	StealReply struct {
		Count int
		Tasks [MAX_TASKS_RPC]nm.MonoidStorage
	}
)

func NewServer(address string, ts chan nm.GoMonoid, sc chan chan nm.MonoidResults) *Server {
	s := &Server{rpc.NewServer(), address, ts, sc}

	if err := s.RPC.RegisterName("Danse", s); err != nil {
		log.Panicln("RPC register failed:", err)
	}

	go func() {
		conn, err := net.Listen("tcp", s.Address)

		if err != nil {
			log.Panicln("Listener accept failed:", err)
		}

		s.RPC.Accept(conn)
	}()

	return s
}

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

func (s *Server) SyncRequest(useless bool, reply *nm.MonoidResults) error {
	ch := make(chan nm.MonoidResults, 1)
	s.Sync <- ch
	*reply = <-ch

	return nil
}
