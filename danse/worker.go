package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"

	kcp "github.com/xtaci/kcp-go"
)

type (
	WorkerState int

	Worker struct {
		Connection net.Conn
		Buffer     bytes.Buffer
	}
)

func NewWorker(address string) (*Worker, error) {
	conn, err := kcp.Dial(address)
	return &Worker{Connection: conn}, err
}

func (w *Worker) SendWork(t *Task) {
	w.Buffer.Reset()

	err := binary.Write(&w.Buffer, binary.BigEndian, *t)

	if err != nil {
		log.Println("Binary write failed:", err.Error())
		return
	}

	_, err = w.Connection.Write(w.Buffer.Bytes())

	if err != nil {
		log.Println("Connection write failed:", err.Error())
	}
}

func (w *Worker) Bye() error {
	return w.Connection.Close()
}
