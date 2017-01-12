package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

const (
	MAX_GENUS = 15
	SIZE      = 3 * (MAX_GENUS - 1)
)

type Monoid struct {
	genus            int
	lastRemovedIndex int
	decs             [SIZE]bool
}

func initDecs() (res [SIZE]bool) {
	res[0] = true

	for i := 1; i < SIZE-1; i++ {
		res[i] = false
	}

	return
}

func explore(m Monoid, c *uint64, d chan bool) {
	m.genus++

	if m.genus > MAX_GENUS {
		d <- true
	}

	for i := 0; i < m.genus; i++ {
		atomic.AddUint64(c, 1)
		go explore(m, c, d)
	}
}

func main() {
	fmt.Printf("Maximum genus: %d\nMaximum procs to be spawned: %d\n", MAX_GENUS, runtime.GOMAXPROCS(0))

	var count uint64 = 0
	var done = make(chan bool)

	explore(Monoid{1, 0, initDecs()}, &count, done)

	if <-done {
		fmt.Printf("Finished computation: %d\n", count)
	}
}
