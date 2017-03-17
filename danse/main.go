package main

import (
	"fmt"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

const (
	MAX_GENUS   = 35
	STACK_BOUND = 11
)

func main() {
	fmt.Printf("Got:\n\t%v\n", WalkChildren(nm.NewMonoid()))
}

func WalkChildren(m nm.GoMonoid) (res nm.MonoidResults) {
	if m.Genus() < MAX_GENUS-STACK_BOUND {
		iter := m.NewIterator()
		rchan := make(chan nm.MonoidResults)
		var nbr uint64 = 0

		// Fork
		go func(it nm.GoGeneratorIterator, ch chan nm.MonoidResults) {
			for it.MoveNext() {
				ch <- WalkChildren(m.RemoveGenerator(it.GetGen()))
				nbr++
			}
			close(ch)
		}(iter, rchan)

		// Join & reduce
		for r := range rchan {
			for k, v := range r {
				res[k] += v
			}
		}

		res[m.Genus()] += nbr
		iter.Free()
	} else {
		m.WalkChildrenStack(&res)
	}

	m.Free()
	return
}
