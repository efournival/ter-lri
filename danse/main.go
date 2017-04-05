package main

import (
	"flag"
	"fmt"

	"github.com/efournival/ter-lri/go-numeric-monoid"
)

const (
	MAX_GENUS   = 35
	STACK_BOUND = 11
)

func main() {
	danser := NewDanser()
	danser.WorkerFunc(func(m nm.GoMonoid) (res nm.MonoidResults) {
		if m.Genus() < MAX_GENUS-STACK_BOUND {
			it := m.NewIterator()
			var srchan []chan nm.MonoidResults
			var nbr uint64 = 0

			for it.MoveNext() {
				srchan = append(srchan, make(chan nm.MonoidResults))
				go func(gen uint64, rchan chan nm.MonoidResults) {
					rchan <- danser.Work(m.RemoveGenerator(gen))
				}(it.GetGen(), srchan[nbr])
				nbr++
			}

			for _, r := range srchan {
				for k, v := range <-r {
					res[k] += v
				}
			}

			res[m.Genus()] += nbr
			it.Free()
		} else {
			m.WalkChildrenStack(&res)
		}

		return
	})

	danser.MasterFunc(func() {
		fmt.Printf("Results for MAX_GENUS=%d:\n\t%v\n", MAX_GENUS, danser.Work(nm.NewMonoid()))
	})

	var master bool
	flag.BoolVar(&master, "master", false, "One to rule them all")
	flag.Parse()

	danser.Danse(master)
}
