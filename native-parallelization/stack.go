package main

import "sync"

type Stack struct {
	sync.RWMutex
	s []*Monoid
}

func NewStack() *Stack {
	return &Stack{s: make([]*Monoid, 0)}
}

func (s *Stack) Push(v *Monoid) {
	s.Lock()
	s.s = append(s.s, v)
	s.Unlock()
}

func (s *Stack) Pop() *Monoid {
	s.Lock()
	l := len(s.s)
	res := s.s[l-1]
	s.s = s.s[:l-1]
	s.Unlock()
	return res
}

func (s *Stack) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()

	return len(s.s) == 0
}
