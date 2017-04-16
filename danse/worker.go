package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"log"
	"net"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

type Worker struct {
	connection net.Conn
	taskStream chan nm.MonoidStorage
	results    chan nm.MonoidResults
}

func NewWorker(address string, ts chan nm.MonoidStorage, r chan nm.MonoidResults) (*Worker, error) {
	conn, err := net.Dial("tcp", address)

	if err != nil {
		return nil, err
	}

	w := &Worker{conn, ts, r}
	go w.waitForAnswers()

	return w, err
}

func (w *Worker) Steal(max int32) error {
	return binary.Write(w.connection, binary.BigEndian, NewStealRequest(max))
}

func (w *Worker) Sync() error {
	return binary.Write(w.connection, binary.BigEndian, NewSyncRequest())
}

func (w *Worker) stealAnswerMessage(sam StealAnswerMessage) {
	for i := 0; i < int(sam.Count); i++ {
		w.taskStream <- sam.Tasks[i]
	}
}

func (w *Worker) syncAnswerMessage(sam SyncAnswerMessage) {
	w.results <- sam.Result
}

func (w *Worker) waitForAnswers() {
	reader := bufio.NewReader(w.connection)

	for {
		b := make([]byte, MAX_MESSAGE_SIZE)
		n, err := reader.Read(b)

		if err != nil {
			log.Println("Reading from", w.connection.LocalAddr(), "failed:", err.Error())
		} else if n > 0 {
			mtype := MessageType(b[0])

			if mtype == StealAnswer {
				var sam StealAnswerMessage

				if err := binary.Read(bytes.NewReader(b), binary.BigEndian, &sam); err != nil {
					log.Println("Binary read (StealAnswer) from", w.connection.LocalAddr(), "failed:", err.Error())
				} else {
					w.stealAnswerMessage(sam)
				}
			} else if mtype == SyncAnswer {
				var sam SyncAnswerMessage

				if err := binary.Read(bytes.NewReader(b), binary.BigEndian, &sam); err != nil {
					log.Println("Binary read (SyncAnswer) from", w.connection.LocalAddr(), "failed:", err.Error())
				} else {
					w.syncAnswerMessage(sam)
				}
			}
		}
	}
}
