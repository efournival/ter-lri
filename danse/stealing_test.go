package main

import (
	"testing"
	"time"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

var (
	server   *Server
	worker   *Worker
	tasks    []nm.MonoidStorage
	tinchan  chan nm.MonoidStorage
	toutchan chan nm.MonoidStorage
)

const (
	ADDR  = "localhost:12345"
	TASKS = 3
	MAX   = 2
)

func init() {
	tinchan = make(chan nm.MonoidStorage, 10)
	toutchan = make(chan nm.MonoidStorage, 10)

	for i := 0; i < TASKS; i++ {
		t := nm.NewMonoid().GetBytes()
		tinchan <- t
		tasks = append(tasks, t)
	}
}

func TestServerInitialization(t *testing.T) {
	t.Log("Number of tasks in input channel is", len(tinchan))

	server = NewServer(ADDR, tinchan)

	go func() {
		err := server.Listen()

		if err != nil {
			t.Fatal(err.Error())
		}
	}()
}

func TestWorkerInitialization(t *testing.T) {
	var err error
	worker, err = NewWorker(ADDR, toutchan)

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestWorkerStealing(t *testing.T) {
	err := worker.Steal(MAX)

	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(100 * time.Millisecond)

	t.Log("Number of tasks in input channel is", len(tinchan))
	t.Log("Number of tasks in output channel is", len(toutchan))

	if len(tinchan) != TASKS-MAX {
		t.Fatal("Stealing failed, expected", TASKS-MAX, "tasks in input channel")
	}

	if len(toutchan) != MAX {
		t.Fatal("Stealing failed, expected", MAX, "tasks in output channel")
	}
}
