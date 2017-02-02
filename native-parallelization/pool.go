package main

import (
	"fmt"
	"sync/atomic"
)

const COUNT_SIZE = MAX_GENUS + 1

type Pool struct {
	stack          *Stack
	activeRoutines int64
	result         [COUNT_SIZE]int64
}

func NewPool() *Pool {
	p := Pool{
		stack:          NewStack(),
		activeRoutines: 0,
	}

	return &p
}

func (p *Pool) Work(s *Monoid) {
	if s.g < MAX_GENUS {
		for x := s.c; x < s.c+s.m; x++ {
			if s.d[x] == 1 {
				p.Add(s.GetSon(x))
			}
		}
	}

	atomic.AddInt64(&p.activeRoutines, -1)
}

func (p *Pool) Add(s *Monoid) {
	atomic.AddInt64(&p.result[s.g], 1)
	p.stack.Push(s)
}

func (p *Pool) Start() {
	p.Add(RootMonoid())

	for p.activeRoutines > 0 || !p.stack.IsEmpty() {
		if !p.stack.IsEmpty() {
			s := p.stack.Pop()
			atomic.AddInt64(&p.activeRoutines, 1)
			go p.Work(s)
		}
	}

	fmt.Printf("%+v\n", p.result)
}
