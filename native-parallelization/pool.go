package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

const COUNT_SIZE = MAX_GENUS + 1

type Pool struct {
	stack  *Stack
	done   chan bool
	result [COUNT_SIZE]int64
}

func NewPool() *Pool {
	p := Pool{}
	p.stack = NewStack()
	p.done = make(chan bool)

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
}

// TODO: received work will be added with this
func (p *Pool) Add(s *Monoid) {
	atomic.AddInt64((*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(&p.result[0]))+(uintptr)(s.g)*unsafe.Sizeof(p.result[0]))), 1)

	if p.result[29] == 3437839 {
		p.done <- true
	}

	p.stack.Push(s)
}

func (p *Pool) Start() {
	p.Add(RootMonoid())

	go func() {
		for {
			if !p.stack.IsEmpty() {
				s := p.stack.Pop()
				go p.Work(s)
			}
		}
	}()

	if <-p.done {
		fmt.Printf("%+v\n", p.result)
	}
}
