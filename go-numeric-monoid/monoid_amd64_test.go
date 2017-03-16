package nm

import "testing"

var first16results = []uint64{1, 2, 4, 7, 12, 23, 39, 67, 118, 204, 343, 592, 1001, 1693, 2857, 4806}

func CheckFirst16Results(results []uint64, t *testing.T) {
	for k, v := range first16results {
		if v != results[k] {
			t.Errorf("Bad result for genus %d, expected %d and got %d\n", k, v, results[k])
			t.Errorf("Expected:\n%+v\n", first16results)
			t.Errorf("Got:\n%+v\n", results[:16])
			t.FailNow()
		}
	}
}

func TestNumericMonoidWalking(t *testing.T) {
	nm := NewMonoid()
	CheckFirst16Results(nm.Walk(), t)
	nm.Free()
}

func TestChildrenStackWalking(t *testing.T) {
	nm := NewMonoid()
	var res MonoidResults
	nm.WalkChildrenStack(&res)
	CheckFirst16Results(res[:], t)
	nm.Free()
}

func TestGenus(t *testing.T) {
	nm := NewMonoid()
	if nm.Genus() != 0 {
		t.Fatalf("The genus of the root monoid should be 0, got %d\n", nm.Genus())
	}
	nm.Free()
}
