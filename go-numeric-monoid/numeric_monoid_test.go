package nm

import "testing"

func Walk(m GoMonoid) (res MonoidResults) {
	if m.Genus() < 35-11 {
		it := m.NewIterator()
		var srchan []chan MonoidResults
		var nbr uint64 = 0

		for it.MoveNext() {
			srchan = append(srchan, make(chan MonoidResults))
			go func(gen uint64, rchan chan MonoidResults) {
				rchan <- Walk(m.RemoveGenerator(gen))
			}(it.GetGen(), srchan[nbr])
			nbr++
		}

		for _, r := range srchan {
			for k, v := range <-r {
				res[k] += v
			}
		}

		res[m.Genus()] += nbr
		it.Free()
	} else {
		m.WalkChildrenStack(&res)
	}

	return
}

func TestNumericMonoid(t *testing.T) {
	t.Logf("Results:\n%+v\n", Walk(NewMonoid()))
}
