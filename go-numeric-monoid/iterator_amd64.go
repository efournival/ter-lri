package nm

// #cgo LDFLAGS: -lcilkrts
// #cgo CXXFLAGS: -std=c++11 -Wno-cpp -fPIC -g -O3 -fcilkplus -march=native -mtune=native
// #include "ctreewalk.h"
import "C"

type GoGeneratorIterator struct {
	gi C.GeneratorIterator
}

func (gm GoMonoid) NewIterator() GoGeneratorIterator {
	var ggi GoGeneratorIterator
	ggi.gi = C.NewGeneratorIterator(gm.m)
	return ggi
}

func (ggi GoGeneratorIterator) MoveNext() bool {
	return C.MoveNext(ggi.gi) != 0
}

// Here, uint == uint_fast64_t == uint64_t
func (ggi GoGeneratorIterator) GetGen() uint {
	return uint(C.GetGen(ggi.gi))
}

func (ggi GoGeneratorIterator) Free() {
	C.FreeGeneratorIterator(ggi.gi)
}
