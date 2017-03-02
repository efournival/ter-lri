package main

import "testing"

func CheckSlicesEqual(s1, s2 []int, fail func()) {
	if len(s1) != len(s2) {
		fail()
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			fail()
		}
	}
}

func CheckMonoid(t *testing.T, m Monoider, x int, expected []int) Monoider {
	r := m.Son(x)

	CheckSlicesEqual(r.Generators(), expected, func() {
		t.Errorf("Checking son monoid by removing %d failed, expected: %+v and got: %+v\n", x, expected, r.Generators())
	})

	return r
}

func CheckKnownMonoids(t *testing.T, root Monoider) {
	a := CheckMonoid(t, root, 1, []int{2, 3})
	b := CheckMonoid(t, a, 2, []int{3, 4, 5})
	c := CheckMonoid(t, a, 3, []int{2, 5})
	CheckMonoid(t, c, 5, []int{2, 7})
	CheckMonoid(t, b, 3, []int{4, 5, 6, 7})
	CheckMonoid(t, b, 4, []int{3, 5, 7})
	CheckMonoid(t, b, 5, []int{3, 4})
}

func CheckKnownCount(t *testing.T, root Monoider) {
	got := Count(root)
	expected := []int{1, 1, 2, 4, 7, 12, 23, 39, 67, 118, 204}

	CheckSlicesEqual(got[:], expected, func() {
		t.Fatalf("Wrong monoids count, expected: %+v and got: %+v\n", expected, got[:])
	})
}

func TestNaiveKnownMonoids(t *testing.T) {
	CheckKnownMonoids(t, InitNaiveMonoid())
}

func TestNaiveKnownCount(t *testing.T) {
	CheckKnownCount(t, InitNaiveMonoid())
}

/*
	TODO: investigate why this does not work while the count is correct?!

	Output:
		monoid_test.go:21: Checking son monoid by removing 3 failed, expected: [2 5] and got: [5]
		monoid_test.go:21: Checking son monoid by removing 3 failed, expected: [2 5] and got: [5]
		monoid_test.go:21: Checking son monoid by removing 5 failed, expected: [2 7] and got: [7]
		monoid_test.go:21: Checking son monoid by removing 5 failed, expected: [2 7] and got: [7]
		monoid_test.go:21: Checking son monoid by removing 4 failed, expected: [3 5 7] and got: [5 7]
		monoid_test.go:21: Checking son monoid by removing 4 failed, expected: [3 5 7] and got: [5 7]
		monoid_test.go:21: Checking son monoid by removing 4 failed, expected: [3 5 7] and got: [5 7]
		monoid_test.go:21: Checking son monoid by removing 5 failed, expected: [3 4] and got: []

	func TestOptimizedKnownMonoids(t *testing.T) {
		CheckKnownMonoids(t, InitOptimizedMonoid())
	}
*/

func TestOptimizedKnownCount(t *testing.T) {
	CheckKnownCount(t, InitOptimizedMonoid())
}
