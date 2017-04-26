package main

import (
	"testing"
	"time"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

func init() {
	for i := 0; i < TASKS; i++ {
		t := nm.NewMonoid()
		in <- t
		tasks = append(tasks, t)
	}
}

func TestWorkerStealing(t *testing.T) {
	t.Log("Number of tasks:", len(in))

	// Wait for worker to connect
	time.Sleep(500 * time.Millisecond)
	worker.Steal()

	t.Log("Number of tasks in input channel:", len(in))
	t.Log("Number of tasks in output channel:", len(out))

	if len(in) != TASKS-STEAL_COUNT {
		t.Fatal("Stealing failed, expected", TASKS-STEAL_COUNT, "tasks in input channel")
	}

	if len(out) != STEAL_COUNT {
		t.Fatal("Stealing failed, expected", STEAL_COUNT, "tasks in output channel")
	}
}
