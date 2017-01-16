package main

import "testing"

func CheckSlicesEqual(t *testing.T, s1, s2 []int, fail func()) {
	if len(s1) != len(s2) {
		fail()
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			fail()
		}
	}
}

func CheckMonoid(t *testing.T, m Monoid, x int, expected []int) Monoid {
	r := son(m, x)
	CheckSlicesEqual(t, r.gen, expected, func() {
		t.Fatalf("Checking son monoid by removing %d failed, expected: %+v and got: %+v\n", x, expected, r.gen)
	})

	return r
}

func TestKnownMonoids(t *testing.T) {
	a := CheckMonoid(t, root(), 1, []int{2, 3})
	b := CheckMonoid(t, a, 2, []int{3, 4, 5})
	c := CheckMonoid(t, a, 3, []int{2, 5})
	CheckMonoid(t, c, 5, []int{2, 7})
	CheckMonoid(t, b, 3, []int{4, 5, 6, 7})
	CheckMonoid(t, b, 4, []int{3, 5, 7})
	CheckMonoid(t, b, 5, []int{3, 4})
}

func TestKnownCount(t *testing.T) {
	got := count()
	expected := []int{1, 1, 2, 4, 7, 12, 23, 39, 67, 118, 204}
	CheckSlicesEqual(t, got[:], expected, func() { t.Fatalf("Wrong monoids count, expected: %+v and got: %+v\n", expected, got[:]) })
}
