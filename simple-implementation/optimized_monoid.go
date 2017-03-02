package main

type OptimizedMonoid struct {
	c int
	g int
	m int
	d [SIZE]int
}

func InitOptimizedMonoid() (m OptimizedMonoid) {
	m.c = 1
	m.g = 0
	m.m = 1

	for i := 0; i < SIZE; i++ {
		m.d[i] = 1 + i/2
	}

	return
}

func (s OptimizedMonoid) Son(x int) Monoider {
	sx := OptimizedMonoid{}
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

	return sx
}

func (s OptimizedMonoid) Generators() (res []int) {
	for x := s.c; x < s.c+s.m; x++ {
		if s.d[x] == 1 {
			res = append(res, x)
		}
	}

	return
}

func (s OptimizedMonoid) Genus() int {
	return s.g
}

func (s OptimizedMonoid) Multiplicity() int {
	return s.m
}
