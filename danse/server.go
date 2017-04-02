package main

import (
	"encoding/binary"
	"log"
	"net"

	kcp "github.com/xtaci/kcp-go"
)

type Server struct {
	Address  string
	Listener net.Listener
	OnAccept func(net.Conn)
}

func NewServer(addr string) (s *Server) {
	s = &Server{addr, nil, nil}

	s.AcceptFunc(func(t Task) {
		log.Println("Receive function is not defined")
	})

	return
}

func (s *Server) Listen() (err error) {
	s.Listener, err = kcp.Listen(s.Address)

	if err != nil {
		return
	}

	for {
		conn, aerr := s.Listener.Accept()

		if aerr != nil {
			log.Println("Listener accept failed:", aerr.Error())
		} else {
			go s.OnAccept(conn)
		}
	}

	return
}

func (s *Server) AcceptFunc(rf func(Task)) {
	s.OnAccept = func(c net.Conn) {
		var tsk Task

		err := binary.Read(c, binary.BigEndian, &tsk)

		if err != nil {
			log.Println("Binary read failed:", err.Error())
			return
		}

		rf(tsk)
	}
}
