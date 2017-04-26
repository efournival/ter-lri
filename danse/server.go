package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"log"
	"net"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

type Server struct {
	Address    string
	TaskStream chan nm.GoMonoid
	Sync       chan net.Conn
}

func NewServer(addr string, ts chan nm.GoMonoid, sc chan net.Conn) (s *Server) {
	s = &Server{addr, ts, sc}
	return
}

func (s *Server) Listen() (err error) {
	var listener net.Listener
	listener, err = net.Listen("tcp", s.Address)

	if err != nil {
		return
	}

	for {
		conn, aerr := listener.Accept()

		if aerr != nil {
			log.Panicln("Listener accept failed:", aerr.Error())
		} else {
			go s.onAccept(conn)
		}
	}

	return
}

func (s *Server) stealRequestMessage(srm StealRequestMessage, conn net.Conn) {
	log.Println("Received steal request from", conn.RemoteAddr())

	var tasks []nm.MonoidStorage

	for i := 0; i < int(srm.Max); i++ {
		select {
		// We can steal a task, add it to our steal answer
		case t := <-s.TaskStream:
			tasks = append(tasks, t.GetBytes())
		// No task left
		// TODO: add timer channel
		default:
			break
		}
	}

	if len(tasks) > 0 {
		err := binary.Write(conn, binary.BigEndian, NewStealAnswer(tasks))

		if err != nil {
			log.Panicln("SERVER Binary write (StealRequest) to", conn.LocalAddr(), "failed:", err.Error())
		}
	}
}

func (s *Server) syncRequestMessage(srm SyncRequestMessage, conn net.Conn) {
	log.Println("Asked to sync from", conn.RemoteAddr())
	s.Sync <- conn
}

func (s *Server) onAccept(conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		b := make([]byte, MAX_MESSAGE_SIZE)
		n, err := reader.Read(b)

		if err != nil {
			log.Panicln("SERVER Reading from", conn.LocalAddr(), "failed:", err.Error())
		} else if n > 0 {
			mtype := MessageType(b[0])

			if mtype == StealRequest {
				var srm StealRequestMessage

				if err := binary.Read(bytes.NewReader(b), binary.BigEndian, &srm); err != nil {
					log.Panicln("SERVER Binary read (StealRequest) from", conn.LocalAddr(), "failed:", err.Error())
				} else {
					s.stealRequestMessage(srm, conn)
				}
			} else if mtype == SyncRequest {
				var srm SyncRequestMessage

				if err := binary.Read(bytes.NewReader(b), binary.BigEndian, &srm); err != nil {
					log.Panicln("SERVER Binary read (SyncRequest) from", conn.LocalAddr(), "failed:", err.Error())
				} else {
					s.syncRequestMessage(srm, conn)
				}
			}
		}
	}
}
