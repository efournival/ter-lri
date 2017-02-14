package nm

import "testing"

func TestNumericMonoid(t *testing.T) {
	nm := NewMonoid()
	res, dur := nm.Walk()

	first16results := []uint{1, 2, 4, 7, 12, 23, 39, 67, 118, 204, 343, 592, 1001, 1693, 2857, 4806}
	for k, v := range first16results {
		if v != res[k] {
			t.Fatalf("Bad result for genus %d, expected %d and got %d\n", k, v, res[k])
		}
	}

	t.Logf("Tree walking took %v\n", dur)

	nm.Free()
}
