package main

import (
	"fmt"
	"testing"
)

const MGENUS = 26

type Result [MGENUS + 1]uint

type OptimizedMonoid struct {
	c int
	g int
	m int
	d [MGENUS * 3]int
}

func InitOptimizedMonoid() (m OptimizedMonoid) {
	m.c = 1
	m.g = 0
	m.m = 1

	for i := 0; i < len(m.d); i++ {
		m.d[i] = 1 + i/2
	}

	return
}

func (s OptimizedMonoid) Son(x int) (sx OptimizedMonoid) {
	sx.c = x + 1
	sx.g = s.g + 1

	if x > s.m {
		sx.m = s.m
	} else {
		sx.m = s.m + 1
	}

	sx.d = s.d

	for y := x; y < len(s.d); y++ {
		if s.d[y-x] > 0 {
			sx.d[y] = s.d[y] - 1
		}
	}

	return
}

var (
	throughputNTasks   chan OptimizedMonoid
	throughputNResults chan int
	throughputNResult  Result
)

func init() {
	throughputNTasks = make(chan OptimizedMonoid, MaxTasks)
	throughputNResults = make(chan int, MaxTasks)
}

func count(s OptimizedMonoid) {
	throughputNResults <- s.g

	if s.g < MGENUS {
		for x := s.c; x < s.c+s.m; x++ {
			if s.d[x] == 1 {
				throughputNTasks <- s.Son(x)
			}
		}
	}
}

func nativeSched() {
	go func() {
		for {
			select {
			case task := <-throughputNTasks:
				count(task)
			case result := <-throughputNResults:
				throughputNResult[result]++
			default:
				finished <- true
				return
			}
		}
	}()
}

func getNTotal(res Result) (accumulator uint) {
	for _, v := range res {
		accumulator += v
	}

	return
}

func testNative(t *testing.T) {
	finished = make(chan bool, 1)
	throughputNTasks <- InitOptimizedMonoid()
	nativeSched()

	if <-finished {
		fmt.Println("Total =", getNTotal(throughputNResult)-1)
	}

	for i := 0; i < len(throughputNResult); i++ {
		throughputNResult[i] = 0
	}
}

func TestNativeThroughput(t *testing.T) {
	for i := 0; i < 4; i++ {
		testNative(t)
	}
}
