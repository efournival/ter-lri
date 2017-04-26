package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"time"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

type Worker struct {
	connection net.Conn
	address    string
	taskStream chan nm.GoMonoid
	results    chan nm.MonoidResults
	lastSync   time.Time
}

func NewWorker(address string, ts chan nm.GoMonoid, r chan nm.MonoidResults) *Worker {
	w := &Worker{nil, address, ts, r, time.Now()}

	go func() {
		for {
			if err := w.Connect(); err == nil {
				log.Println("Successfully connected to", address)
				return
			}

			time.Sleep(250 * time.Millisecond)
		}
	}()

	return w
}

func (w *Worker) Connect() error {
	conn, err := net.Dial("tcp", w.address)

	if err != nil {
		return err
	}

	w.connection = conn
	go w.waitForAnswers()

	return nil
}

func (w *Worker) Steal(max int32) error {
	return binary.Write(w.connection, binary.BigEndian, NewStealRequest(max))
}

func (w *Worker) Sync() error {
	return binary.Write(w.connection, binary.BigEndian, NewSyncRequest())
}

func (w *Worker) stealAnswerMessage(sam StealAnswerMessage) {
	w.lastSync = time.Now()

	log.Println("Stole", sam.Count, "tasks")

	for i := 0; i < int(sam.Count); i++ {
		w.taskStream <- nm.FromBytes(sam.Tasks[i])
	}
}

func (w *Worker) syncAnswerMessage(sam SyncAnswerMessage) {
	log.Println("Received sync from", w.connection.RemoteAddr())
	w.results <- sam.Result
}

func (w *Worker) waitForAnswers() {
	reader := bufio.NewReader(w.connection)

	for {
		b := make([]byte, MAX_MESSAGE_SIZE)
		n, err := reader.Read(b)

		if err != nil {
			log.Panicln("WORKER Reading from", w.connection.LocalAddr(), "failed:", err.Error())
		} else if n > 0 {
			mtype := MessageType(b[0])

			if mtype == StealAnswer {
				var sam StealAnswerMessage

				if err := binary.Read(bytes.NewReader(b), binary.BigEndian, &sam); err != nil {
					log.Panicln("WORKER Binary read (StealAnswer) from", w.connection.LocalAddr(), "failed:", err.Error())
				} else {
					w.stealAnswerMessage(sam)
				}
			} else if mtype == SyncAnswer {
				var sam SyncAnswerMessage

				if err := binary.Read(bytes.NewReader(b), binary.BigEndian, &sam); err != nil {
					log.Panicln("WORKER Binary read (SyncAnswer) from", w.connection.LocalAddr(), "failed:", err.Error())
				} else {
					w.syncAnswerMessage(sam)
				}
			}
		}
	}
}
