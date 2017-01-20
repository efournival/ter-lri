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
	// x MUST BE a generator of the source monoid
	isGenerator := false

	for _, v := range s.gen {
		if v < x {
			// We only need generators before x to compute this new monoid
			sx.gen = append(sx.gen, v)
		} else if v == x {
			// Little check for more safety
			isGenerator = true
			// Break is OK because numbers are in the ascending order
			break
		}
	}

	if !isGenerator {
		panic(fmt.Sprintf("ERROR: %d is not a generator\n", x))
	}

	sx.g = s.g + 1
	sx.d = s.d
	// Remove x
	sx.d[x] = false
	sx.m = x

	// We want to find new generators AFTER the generator we just removed
	for g := x + 1; g < SIZE; g++ {
		// If the number is in the semigroup
		if sx.d[g] {
			found := false

			// We are looking for non-generator numbers
			for i := 1; i <= g/2; i++ {
				if sx.d[i] && sx.d[g-i] {
					// g is a decomposition number, not a generator
					found = true
					break
				}
			}

			if !found {
				// If g is not a decomposition number, we have found a generator
				sx.gen = append(sx.gen, g)
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
				if x >= s.m {
					stack.Push(son(s, x))
				}
			}
		}
	}

	return
}

func main() {
	fmt.Printf("%+v\n", count())
}
