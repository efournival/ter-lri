package main

type Stack struct {
	s []Monoider
}

func NewStack() *Stack {
	return &Stack{s: make([]Monoider, 0)}
}

func (s *Stack) Push(v Monoider) {
	s.s = append(s.s, v)
}

func (s *Stack) Pop() Monoider {
	l := len(s.s)
	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res
}

func (s *Stack) IsEmpty() bool {
	return len(s.s) == 0
}
