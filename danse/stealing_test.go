package main

import (
	"testing"
	"time"

	nm "github.com/efournival/ter-lri/go-numeric-monoid"
)

func init() {
	for i := 0; i < Tasks; i++ {
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

	if len(in) != Tasks-MaxTasksRPC {
		t.Fatal("Stealing failed, expected", Tasks-MaxTasksRPC, "tasks in input channel")
	}

	if len(out) != MaxTasksRPC {
		t.Fatal("Stealing failed, expected", MaxTasksRPC, "tasks in output channel")
	}
}
