package main

import (
	"bytes"
	"encoding/binary"
	"net"

	kcp "github.com/xtaci/kcp-go"
)

type (
	WorkerState int

	worker struct {
		connection net.Conn
		buffer     bytes.Buffer
	}
)

func NewWorker(address string) (*worker, error) {
	conn, err := kcp.Dial(address)
	return &worker{connection: conn}, err
}

func (w *worker) SendWork(t *task) {
	w.buffer.Reset()
	binary.Write(&w.buffer, binary.BigEndian, t)
	w.connection.Write(w.buffer.Bytes())
}

func (w *worker) Bye() error {
	return w.connection.Close()
}
