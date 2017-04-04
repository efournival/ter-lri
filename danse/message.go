package main

import (
	"sync/atomic"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

const MAX_TASKS_IN_MESSAGE = 10

type (
	StealRequestMessage struct {
		Max int32
	}

	StealAnswerMessage struct {
		Count int32
		Tasks [MAX_TASKS_IN_MESSAGE]Task
	}

	Task struct {
		Identifier uint64
		Data       nm.MonoidStorage
		Results    nm.MonoidResults
	}
)

var incTask uint64

func NewTask(data nm.GoMonoid) (res *Task) {
	var mr nm.MonoidResults
	res = &Task{incTask, data.GetBytes(), mr}

	// TODO: mutex?
	atomic.AddUint64(&incTask, 1)

	return
}

func NewStealRequest(max int32) StealRequestMessage {
	return StealRequestMessage{max}
}

func NewStealAnswer(tasks []Task) (res StealAnswerMessage) {
	res.Count = int32(len(tasks))
	copy(res.Tasks[:], tasks)
	return
}
