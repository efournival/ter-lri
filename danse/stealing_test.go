package main

import (
	"testing"
	"time"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

func init() {
	tinchan = make(chan nm.GoMonoid, 10)
	toutchan = make(chan nm.GoMonoid, 10)

	for i := 0; i < TASKS; i++ {
		t := nm.NewMonoid()
		tinchan <- t
		tasks = append(tasks, t)
	}
}

func TestWorkerStealing(t *testing.T) {
	t.Log("Number of tasks:", len(tinchan))

	// Wait for worker to connect
	time.Sleep(500 * time.Millisecond)
	err := worker.Steal(MAX)

	if err != nil {
		t.Fatal(err)
	}

	// Wait for packet to be transmitted
	time.Sleep(100 * time.Millisecond)

	t.Log("Number of tasks in input channel:", len(tinchan))
	t.Log("Number of tasks in output channel:", len(toutchan))

	if len(tinchan) != TASKS-MAX {
		t.Fatal("Stealing failed, expected", TASKS-MAX, "tasks in input channel")
	}

	if len(toutchan) != MAX {
		t.Fatal("Stealing failed, expected", MAX, "tasks in output channel")
	}
}
