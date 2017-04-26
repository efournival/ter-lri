package main

import "github.com/efournival/ter-lri/go-numeric-monoid"

var (
	server  *Server
	worker  *Worker
	tasks   []nm.GoMonoid
	in, out chan nm.GoMonoid
	syncc   chan chan nm.MonoidResults
	results chan nm.MonoidResults
)

const (
	ADDR  = "localhost:12345"
	TASKS = 1000
)

func init() {
	in = make(chan nm.GoMonoid, TASKS)
	out = make(chan nm.GoMonoid, TASKS)
	syncc = make(chan chan nm.MonoidResults, 1)
	results = make(chan nm.MonoidResults, MAX_TASKS)

	server = NewServer(ADDR, in, syncc)
	worker = NewWorker(ADDR, out, results)
}
