package main

import (
	"encoding/binary"
	"net"

	kcp "github.com/xtaci/kcp-go"
)

type (
	ReceiveFunc func(net.Conn)

	server struct {
		address     string
		listener    net.Listener
		receiveFunc ReceiveFunc
	}
)

func NewServer(port string) *server {
	return &server{"localhost:" + port, nil, nil}
}

func (s *server) Listen(finished chan bool) (err error) {
	s.listener, err = kcp.Listen(s.address)

	if err != nil {
		return
	}

	go func() {
		for {
			conn, err := s.listener.Accept()

			if err == nil {
				go s.receiveFunc(conn)
			}
		}

		close(finished)
	}()

	return
}

func (s *server) ReceiveFunc(rf func(task)) {
	s.receiveFunc = func(c net.Conn) {
		var tsk task
		if binary.Read(c, binary.BigEndian, &tsk) == nil {
			rf(tsk)
		}
	}
}
