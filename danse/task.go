package main

import "github.com/efournival/ter-lri/go-numeric-monoid"

type (
	TaskState uint8

	task struct {
		identifier uint64
		state      TaskState
		data       nm.GoMonoid
		results    nm.MonoidResults
	}
)

const (
	Waiting TaskState = iota
	Processing
)

func NewTask(data nm.GoMonoid, identifier uint64) *task {
	var mr nm.MonoidResults
	return &task{identifier, Waiting, data, mr}
}
