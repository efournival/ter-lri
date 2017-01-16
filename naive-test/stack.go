package main

type stack struct {
	s []Monoid
}

func NewStack() *stack {
	return &stack{s: make([]Monoid, 0)}
}

func (s *stack) Push(v Monoid) {
	s.s = append(s.s, v)
}

func (s *stack) Pop() Monoid {
	l := len(s.s)
	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res
}

func (s *stack) IsEmpty() bool {
	return len(s.s) == 0
}
