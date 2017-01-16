package main

import "fmt"

const (
	MAX_GENUS  = 10
	COUNT_SIZE = MAX_GENUS + 1
	SIZE       = 3 * MAX_GENUS
)

type Monoid struct {
	c int
	g int
	m int
	d [SIZE]int
}

func root() Monoid {
	// Paper say c=0 but it's not working, neither does changing the count loop < to <=
	m := Monoid{c: 1, g: 0, m: 1}

	for i := 0; i < SIZE; i++ {
		m.d[i] = 1 + i/2
	}

	return m
}

func son(s Monoid, x int) (sx Monoid) {
	sx.c = x + 1
	sx.g = s.g + 1

	if x > s.m {
		sx.m = s.m
	} else {
		sx.m = s.m + 1
	}

	sx.d = s.d

	for y := x; y < SIZE; y++ {
		if s.d[y-x] > 0 {
			sx.d[y] = s.d[y] - 1
		}
	}

	return
}

func count() (n [COUNT_SIZE]int) {
	stack := NewStack()
	stack.Push(root())

	for !stack.IsEmpty() {
		s := stack.Pop()

		n[s.g]++

		if s.g < MAX_GENUS {
			// < or <= ?
			for x := s.c; x < s.c+s.m; x++ {
				if s.d[x] == 1 {
					stack.Push(son(s, x))
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
