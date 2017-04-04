package main

import (
	"encoding/binary"
	"log"
	"net"

	kcp "github.com/xtaci/kcp-go"
)

type Server struct {
	Address    string
	TaskStream chan Task
}

func NewServer(addr string, ts chan Task) (s *Server) {
	s = &Server{addr, ts}
	return
}

func (s *Server) Listen() (err error) {
	var listener net.Listener
	listener, err = kcp.Listen(s.Address)

	if err != nil {
		return
	}

	for {
		conn, aerr := listener.Accept()

		if aerr != nil {
			log.Println("Listener accept failed:", aerr.Error())
		} else {
			go s.onAccept(conn)
		}
	}

	return
}

func (s *Server) onAccept(conn net.Conn) {
	var srm StealRequestMessage
	err := binary.Read(conn, binary.BigEndian, &srm)

	if err != nil {
		log.Println("Binary read failed:", err.Error())
		return
	}

	// TODO: min value
	if len(s.TaskStream) > 0 {
		var tasks []Task

		for i := 0; i < int(srm.Max); i++ {
			select {
			// We can steal a task, add it to our steal answer
			case t := <-s.TaskStream:
				tasks = append(tasks, t)
			// No task left
			default:
				break
			}
		}

		err = binary.Write(conn, binary.BigEndian, NewStealAnswer(tasks))

		if err != nil {
			log.Println("Binary write failed:", err.Error())
		}
	}
}
