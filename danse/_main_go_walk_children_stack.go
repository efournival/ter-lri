package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

const (
	MAX_GENUS   = 40
	STACK_BOUND = 11
)

var r nm.MonoidResults

func main() {
	m := nm.NewMonoid()
	WalkChildrenStack(m)
	//m.WalkChildrenStack(&r)	// utilise le binding C++ (fonctionne)
	fmt.Printf("Got:\n\t%v\n", r)
}

func WalkChildren(m nm.GoMonoid) {
	if m.Genus() < MAX_GENUS-STACK_BOUND {
		var wg sync.WaitGroup
		var nbr uint64 = 0
		it := m.NewIterator()

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
		WalkChildrenStack(m)
	}

	m.Free()
}

func WalkChildrenStack(m nm.GoMonoid) {
	log.Println("Entering WalkChildrenStack")

	var nbr uint64 = 0
	var stack [MAX_GENUS]nm.GoMonoid
	stack_pointer := 1

	stack[0] = m
	//for i := 1; i < MAX_GENUS; i++ {
	//	stack[i] = nm.NewMonoid() // init full n
	//}

	log.Println("Stack initialized")

	log.Println("First element of the stack set to m")

	for stack_pointer != 0 {
		stack_pointer--          // sp = m
		current := stack_pointer // current = m (0)
		log.Printf("Stack pointer = %d, getting iterator\n", stack_pointer)

		it := stack[current].NewIterator()

		log.Println("Got iterator")

		if stack[current].Genus() < MAX_GENUS-1 {
			log.Println("Current genus is smaller than MAX_GENUS-1")

			for it.MoveNext() {
				log.Println("Move next")
				stack_pointer++ // sp = 1
				log.Printf("Top stack is now top+1, about to remove %d\n", it.GetGen())
				stack[stack_pointer] = stack[current].RemoveGenerator(it.GetGen())
				log.Println("Removed generator")
				//stack_pointer++
				nbr++
			}

			log.Println("Resetting stack pointer")
			stack_pointer = current
			atomic.AddUint64(&r[stack[current].Genus()], nbr)
		} else {
			log.Println("Current genus is >= MAX_GENUS-1")
			atomic.AddUint64(&r[stack[current].Genus()], uint64(it.Count()))
		}
	}
}
