package main

import nm "github.com/efournival/ter-lri/go-numeric-monoid"

var (
	server   *Server
	worker   *Worker
	tasks    []nm.GoMonoid
	in, out  chan nm.GoMonoid
	syncc    chan chan nm.MonoidResults
	results  chan nm.MonoidResults
	finished chan bool
)

const (
	Address = "localhost:12345"
	Tasks   = 1000
)

func init() {
	in = make(chan nm.GoMonoid, Tasks)
	out = make(chan nm.GoMonoid, Tasks)
	syncc = make(chan chan nm.MonoidResults, 1)
	results = make(chan nm.MonoidResults, MaxTasks)

	server = NewServer(Address, in, syncc)
	worker = NewWorker(Address, out, results)
}
