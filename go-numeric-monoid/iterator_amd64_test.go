package nm

import "testing"

func TestIterator(t *testing.T) {
	i := NewMonoid().NewIterator()
	var out []uint

	for i.MoveNext() {
		out = append(out, i.GetGen())
	}

	if len(out) != 1 || out[0] != 1 {
		t.Fatalf("Expected [1] and got %+v\n", out)
	}
}
