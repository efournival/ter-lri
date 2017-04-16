package main

import (
	"net"
	"testing"
	"time"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

func init() {
	syncchan = make(chan net.Conn, 1)
	reschan = make(chan nm.MonoidResults, MAX_TASKS)

	var result nm.MonoidResults
	for i := 0; i < MAX_GENUS; i++ {
		result[i] = uint64(i)
	}

	// Worker danser will receive this
	go func() {
		for {
			sync(<-syncchan, result)
		}
	}()
}

func TestWorkerSync(t *testing.T) {
	err := worker.Sync()

	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(100 * time.Millisecond)

	t.Log("Number of results in channel is", len(reschan))

	r := <-reschan
	equals := true

	for i := 0; i < MAX_GENUS; i++ {
		equals = equals && r[i] == uint64(i)
	}

	if !equals {
		t.Fatal("Received result is garbage")
	}
}
