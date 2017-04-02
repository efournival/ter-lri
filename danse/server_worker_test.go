package main

import (
	"testing"
	"time"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

var (
	server       *Server
	worker       *Worker
	receivedTask Task
)

const PORT = ":12345"

func TestServerInitialization(t *testing.T) {
	server = NewServer(PORT)

	server.AcceptFunc(func(tsk Task) {
		receivedTask = tsk
	})

	go func() {
		err := server.Listen()

		if err != nil {
			t.Fatal(err.Error())
		}
	}()
}

func TestWorkerInitialization(t *testing.T) {
	var err error
	worker, err = NewWorker(PORT)

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestWorkerSendWork(t *testing.T) {
	worker.SendWork(NewTask(nm.NewMonoid(), 1))
	time.Sleep(100 * time.Millisecond)

	if *NewTask(nm.NewMonoid(), 1) != receivedTask {
		t.Fatal("Received and sent tasks differ")
	}
}
