package fj

import "testing"

func TestWikipediaTrivialExample(tt *testing.T) {
	g := func(w WorkType) WorkType {
		return w.(int) * 2
	}

	h := func(a int) int {
		t := NewLocalTasks()
		t.Fork(g, a)
		c := a + 1
		t.Join()
		b := t.Result(0).(int)
		return b + c
	}

	f := func(a, b int) int {
		t := NewLocalTasks()
		t.Fork(g, a)
		d := h(b)
		t.Join()
		c := t.Result(0).(int)
		return c + d
	}

	if f(1, 2) != 9 {
		tt.Fatalf("f(1, 2) = %d != 7\n", f(1, 2))
	}
}
