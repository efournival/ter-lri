package main

import (
	"testing"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

func TestNewTask(t *testing.T) {
	nm := nm.NewMonoid()
	tsk := NewTask(nm)

	if len(tsk.Data) != len(nm.GetBytes()) {
		t.Fatal("Task is badly initialized")
	} else {
		t.Log("Task identifier is", tsk.Identifier)
	}
}
