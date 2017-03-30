package main

import (
	"testing"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

func TestNewTask(t *testing.T) {
	nm := nm.NewMonoid()
	tsk := NewTask(nm, 0)

	if tsk.state != Waiting || tsk.data != nm {
		t.Fatal("Task is badly initialized")
	}

	nm.Free()
}
