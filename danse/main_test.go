package main

import (
	"testing"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

var first16results = []uint64{1, 2, 4, 7, 12, 23, 39, 67, 118, 204, 343, 592, 1001, 1693, 2857, 4806}

func CheckFirst16Results(results nm.MonoidResults, t *testing.T) {
	for k, v := range first16results {
		if v != results[k] {
			t.Errorf("Bad result for genus %d, expected %d and got %d\n", k, v, results[k])
			t.Errorf("Expected:\n%+v\n", first16results)
			t.Errorf("Got:\n%+v\n", results[:16])
			t.FailNow()
		}
	}
}

func TestWalkChildren(t *testing.T) {
	CheckFirst16Results(WalkChildren(nm.NewMonoid()), t)
}
