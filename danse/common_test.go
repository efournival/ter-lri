package main

import (
	"net"
	"testing"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

var (
	server   *Server
	worker   *Worker
	tasks    []nm.MonoidStorage
	tinchan  chan nm.MonoidStorage
	toutchan chan nm.MonoidStorage
	syncchan chan net.Conn
	reschan  chan nm.MonoidResults
)

const (
	ADDR  = "localhost:12345"
	TASKS = 3
	MAX   = 2
)

func TestServerInitialization(t *testing.T) {
	t.Log("Number of tasks in input channel is", len(tinchan))

	server = NewServer(ADDR, tinchan, syncchan)

	go func() {
		err := server.Listen()

		if err != nil {
			t.Fatal(err.Error())
		}
	}()
}

func TestWorkerInitialization(t *testing.T) {
	var err error
	worker, err = NewWorker(ADDR, toutchan, reschan)

	if err != nil {
		t.Fatal(err.Error())
	}
}
