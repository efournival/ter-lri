package nm

/*
	#cgo LDFLAGS: -lcilkrts -Xlinker --allow-multiple-definition
	#cgo CXXFLAGS: -std=c++11 -Wno-cpp -fPIC -g -O3 -fcilkplus -march=native -mtune=native
	#include "ctreewalk.h"
	#include <stdlib.h>
*/
import "C"
import "unsafe"

type GoGeneratorIterator struct {
	gi C.GeneratorIterator
}

func (gm GoMonoid) NewIterator() GoGeneratorIterator {
	var ggi GoGeneratorIterator
	ggi.gi = C.NewGeneratorIterator(gm.m)

	// This is bugged for an unknown reason
	//runtime.SetFinalizer(&ggi, func(ggi *GoGeneratorIterator) {
	//	C.free(unsafe.Pointer(ggi.gi))
	//})

	return ggi
}

func (ggi GoGeneratorIterator) MoveNext() bool {
	return C.MoveNext(ggi.gi) != 0
}

func (ggi GoGeneratorIterator) GetGen() uint64 {
	return uint64(C.GetGen(ggi.gi))
}

func (ggi GoGeneratorIterator) Count() uint8 {
	return uint8(C.Count(ggi.gi))
}

func (ggi GoGeneratorIterator) Free() {
	C.free(unsafe.Pointer(ggi.gi))
}
