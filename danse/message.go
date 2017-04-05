package main

import "github.com/efournival/ter-lri/go-numeric-monoid"

const MAX_TASKS_IN_MESSAGE = 10

type (
	StealRequestMessage struct {
		Max int32
	}

	StealAnswerMessage struct {
		Count int32
		Tasks [MAX_TASKS_IN_MESSAGE]nm.MonoidStorage
	}
)

func NewStealRequest(max int32) StealRequestMessage {
	return StealRequestMessage{max}
}

func NewStealAnswer(tasks []nm.MonoidStorage) (res StealAnswerMessage) {
	res.Count = int32(len(tasks))
	copy(res.Tasks[:], tasks)
	return
}
