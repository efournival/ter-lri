package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

const (
	MAX_GENUS   = 35
	STACK_BOUND = 11
)

var r nm.MonoidResults

func main() {
	m := nm.NewMonoid()
	WalkChildren(m)
	fmt.Printf("Got:\n\t%v\n", r)
}

func WalkChildren(m nm.GoMonoid) {
	var wg sync.WaitGroup

	if m.Genus() < MAX_GENUS-STACK_BOUND {
		it := m.NewIterator()
		var nbr uint64 = 0

		for it.MoveNext() {
			wg.Add(1)
			go func(gen uint) {
				WalkChildren(m.RemoveGenerator(gen))
				wg.Done()
			}(it.GetGen())
			nbr++
		}

		wg.Wait()
		atomic.AddUint64(&r[m.Genus()], nbr)
	} else {
		m.WalkChildrenStack(&r)
	}

	m.Free()
}
