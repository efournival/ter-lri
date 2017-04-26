package main

import (
	"net"
	"testing"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

var (
	server   *Server
	worker   *Worker
	tasks    []nm.GoMonoid
	tinchan  chan nm.GoMonoid
	toutchan chan nm.GoMonoid
	syncchan chan net.Conn
	reschan  chan nm.MonoidResults
)

const (
	ADDR  = "localhost:12345"
	TASKS = 3
	MAX   = 2
)

func TestNetworkInitialization(t *testing.T) {
	server = NewServer(ADDR, tinchan, syncchan)

	go func() {
		err := server.Listen()

		if err != nil {
			t.Fatal(err.Error())
		}
	}()

	worker = NewWorker(ADDR, toutchan, reschan)
}
