package main

import (
	"testing"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

func init() {
	var result nm.MonoidResults
	for i := 0; i < MAX_GENUS; i++ {
		result[i] = uint64(i)
	}

	// Worker danser will receive this
	go func() {
		for {
			sync(<-syncc, &result)
		}
	}()
}

func TestWorkerSync(t *testing.T) {
	worker.Sync()

	r := <-results
	equals := true

	for i := 0; i < MAX_GENUS; i++ {
		equals = equals && r[i] == uint64(i)
	}

	if !equals {
		t.Fatal("Received result is garbage")
	}
}
