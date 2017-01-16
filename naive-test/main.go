package main

import "fmt"

const (
	MAX_GENUS  = 10
	COUNT_SIZE = MAX_GENUS + 1
	SIZE       = 3 * MAX_GENUS
)

type Monoid struct {
	g   int
	m   int
	d   [SIZE]bool
	gen []int
}

func root() Monoid {
	m := Monoid{g: 0, m: 1}

	for i := 0; i < SIZE; i++ {
		m.d[i] = true
	}

	m.gen = append(m.gen, 1)

	return m
}

func son(s Monoid, x int) (sx Monoid) {
	sx.g = s.g + 1
	sx.d = s.d
	sx.d[x] = false
	sx.m = x

	for g := s.m; g < SIZE; g++ {
		if sx.d[g] {
			for i := 1; i < (g-1)/2; i++ {
				if sx.d[i] && sx.d[g-1] {
					sx.gen = append(sx.gen, g)
				}
			}
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
			for _, x := range s.gen {
				stack.Push(son(s, x))
			}
		}
	}

	return
}

func main() {
	fmt.Printf("%+v\n", count())
}
