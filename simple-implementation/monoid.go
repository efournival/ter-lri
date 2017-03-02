package main

type Monoider interface {
	Son(x int) Monoider
	Generators() []int
	Genus() int
	Multiplicity() int
}

func Count(root Monoider) (n [COUNT_SIZE]int) {
	stack := NewStack()
	stack.Push(root)

	for !stack.IsEmpty() {
		s := stack.Pop()

		n[s.Genus()]++

		if s.Genus() < MAX_GENUS {
			for _, x := range s.Generators() {
				if x >= s.Multiplicity() {
					stack.Push(s.Son(x))
				}
			}
		}
	}

	return
}
