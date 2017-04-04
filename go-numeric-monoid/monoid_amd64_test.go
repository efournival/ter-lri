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

func TestWalkChildren(t *testing.T) {
	nm := NewMonoid()
	res := nm.WalkChildren()
	CheckFirst16Results(res[:], t)
}

func TestWalkChildrenStack(t *testing.T) {
	nm := NewMonoid()
	var res MonoidResults
	nm.WalkChildrenStack(&res)
	CheckFirst16Results(res[:], t)
}

func TestGenus(t *testing.T) {
	nm := NewMonoid()

	if nm.Genus() != 0 {
		t.Fatalf("The genus of the root monoid should be 0, got %d\n", nm.Genus())
	}
}

func TestGetBytes(t *testing.T) {
	nm := NewMonoid()
	bytes := nm.GetBytes()
	t.Logf("Got %d bytes\n", len(bytes))
}

func TestBytesCopy(t *testing.T) {
	nm1 := NewMonoid()
	bytes1 := nm1.GetBytes()
	nm2 := FromBytes(bytes1)
	bytes2 := nm2.GetBytes()

	for i, b := range bytes1 {
		if b != bytes2[i] {
			t.Fatalf("After bytes copy, monoids differ (%d)", i)
		}
	}
}
