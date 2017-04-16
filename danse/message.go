package main

import "github.com/efournival/ter-lri/go-numeric-monoid"

type (
	MessageType byte

	Message struct {
		Type MessageType
	}

	StealRequestMessage struct {
		Message
		Max int32
	}

	StealAnswerMessage struct {
		Message
		Count int32
		Tasks [MAX_TASKS_IN_MESSAGE]nm.MonoidStorage
	}

	SyncRequestMessage struct {
		Message
	}

	SyncAnswerMessage struct {
		Message
		Result nm.MonoidResults
	}
)

const (
	MAX_TASKS_IN_MESSAGE = 10
	MAX_MESSAGE_SIZE     = 5000

	StealRequest MessageType = iota
	StealAnswer
	SyncRequest
	SyncAnswer
)

func NewStealRequest(max int32) (res StealRequestMessage) {
	res.Type = StealRequest
	res.Max = max
	return
}

func NewStealAnswer(tasks []nm.MonoidStorage) (res StealAnswerMessage) {
	res.Type = StealAnswer
	res.Count = int32(len(tasks))
	copy(res.Tasks[:], tasks)
	return
}

func NewSyncRequest() (res SyncRequestMessage) {
	res.Type = SyncRequest
	return
}

func NewSyncAnswer(result nm.MonoidResults) (res SyncAnswerMessage) {
	res.Type = SyncAnswer
	res.Result = result
	return
}
