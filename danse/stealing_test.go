package main

import (
	"testing"
	"time"

	"github.com/efournival/ter-lri/go-numeric-monoid"
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
