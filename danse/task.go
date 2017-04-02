package main

import "github.com/efournival/ter-lri/go-numeric-monoid"

type (
	TaskState uint8

	Task struct {
		Identifier uint64
		State      TaskState
		Data       nm.MonoidStorage
		Results    nm.MonoidResults
	}
)

const (
	Waiting TaskState = iota
	Processing
)

func NewTask(data nm.GoMonoid, identifier uint64) *Task {
	var mr nm.MonoidResults
	return &Task{identifier, Waiting, data.GetBytes(), mr}
}
