package main

import "fmt"

type NaiveMonoid struct {
	g   int
	m   int
	d   [SIZE]bool
	gen []int
}

func InitNaiveMonoid() (m NaiveMonoid) {
	m.g = 0
	m.m = 1

	for i := 0; i < SIZE; i++ {
		m.d[i] = true
	}

	m.gen = append(m.gen, 1)

	return
}

func (s NaiveMonoid) Son(x int) Monoider {
	sx := NaiveMonoid{}

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

	return sx
}

func (m NaiveMonoid) Generators() []int {
	return m.gen
}

func (m NaiveMonoid) Genus() int {
	return m.g
}

func (m NaiveMonoid) Multiplicity() int {
	return m.m
}
