package main

import "fmt"

func count() (n [COUNT_SIZE]int) {
	stack := NewStack()
	stack.Push(RootMonoid())

	for !stack.IsEmpty() {
		s := stack.Pop()

		n[s.g]++

		if s.g < MAX_GENUS {
			// < or <= ?
			for x := s.c; x < s.c+s.m; x++ {
				if s.d[x] == 1 {
					stack.Push(s.GetSon(x))
				}
			}
		}
	}

	return
}

func main() {
	//fmt.Printf("Maximum genus: %d\nMaximum procs to be spawned: %d\n", MAX_GENUS, runtime.GOMAXPROCS(0))
	fmt.Printf("%+v\n", count())
}
