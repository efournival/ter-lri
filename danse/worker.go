package main

import (
	"encoding/binary"
	"net"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

type Worker struct {
	connection net.Conn
	taskStream chan nm.MonoidStorage
}

func NewWorker(address string, ts chan nm.MonoidStorage) (*Worker, error) {
	conn, err := net.Dial("tcp", address)

	if err != nil {
		return nil, err
	}

	w := &Worker{conn, ts}
	go w.waitForAnswers()

	return w, err
}

func (w *Worker) Steal(max int32) error {
	return binary.Write(w.connection, binary.BigEndian, NewStealRequest(max))
}

func (w *Worker) waitForAnswers() {
	for {
		var sam StealAnswerMessage
		binary.Read(w.connection, binary.BigEndian, &sam)

		for i := 0; i < int(sam.Count); i++ {
			w.taskStream <- sam.Tasks[i]
		}
	}
}
