package main

import (
	"testing"

	nm "github.com/efournival/ter-lri/go-numeric-monoid"
)

var (
	throughputCTasks   chan nm.GoMonoid
	throughputCResults chan nm.MonoidResults
	throughputCResult  nm.MonoidResults
	cilkBound          uint64
)

func init() {
	throughputCTasks = make(chan nm.GoMonoid, MaxTasks)
	throughputCResults = make(chan nm.MonoidResults, MaxTasks)
}

func work(m nm.GoMonoid) {
	if m.Genus() < MaxGenus-cilkBound {
		it := m.NewIterator()
		var nbr uint64

		for it.MoveNext() {
			throughputCTasks <- m.RemoveGenerator(it.GetGen())
			nbr++
		}

		var res nm.MonoidResults
		res[m.Genus()] = nbr
		throughputCResults <- res

		it.Free()
	} else {
		throughputCResults <- m.WalkChildren()
	}

	m.Free()
}

func sched() {
	go func() {
		for {
			select {
			case task := <-throughputCTasks:
				work(task)
			case result := <-throughputCResults:
				for k, v := range result {
					throughputCResult[k] += v
				}
			default:
				finished <- true
				return
			}
		}
	}()
}

func getTotal(res nm.MonoidResults) (accumulator uint64) {
	for _, v := range res {
		accumulator += v
	}

	return
}

func testCilkBound(cb uint64, t *testing.T) {
	finished = make(chan bool, 1)
	cilkBound = cb
	work(nm.NewMonoid())
	sched()

	if <-finished {
		t.Log("CilkBound =", cb, "Total =", getTotal(throughputCResult))
	}

	for i := 0; i < len(throughputCResult); i++ {
		throughputCResult[i] = 0
	}
}

func TestCgoThroughput(t *testing.T) {
	for _, cb := range []uint64{14, 15, 16, 20, 25} {
		testCilkBound(cb, t)
	}
}
