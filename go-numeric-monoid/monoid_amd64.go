package nm

/*
	#cgo LDFLAGS: -lcilkrts -Xlinker --allow-multiple-definition
	#cgo CXXFLAGS: -std=c++11 -Wno-cpp -fPIC -g -O3 -fcilkplus -march=native -mtune=native
	#include "ctreewalk.h"
	#include <stdlib.h>
*/
import "C"

import (
	"unsafe"

	"github.com/intel-go/cpuid"
)

type (
	MonoidResults [C.MAX_GENUS]uint64
	MonoidStorage []byte

	GoMonoid struct {
		m            C.Monoid
		noNeedToFree bool
	}
)

func init() {
	if C.MAX_GENUS < 16 {
		panic("This package has been compiled with MAX_GENUS <= 16 but it is not supported")
	}

	if !cpuid.HasFeature(cpuid.SSSE3) {
		panic("This program requires the SSE3 instructions set")
	}

	if !cpuid.HasFeature(cpuid.POPCNT) {
		panic("This program requires the popcount instruction")
	}

	if int(C.InitCilk()) <= 0 {
		panic("Cilk initialization failed")
	}
}

func NewMonoid() GoMonoid {
	var gm GoMonoid
	gm.m = C.InitFullN()
	gm.noNeedToFree = false
	return gm
}

func FromBytes(ms MonoidStorage) GoMonoid {
	var gm GoMonoid
	gm.m = C.Monoid(&ms[0])
	gm.noNeedToFree = true
	return gm
}

func (gm GoMonoid) Free() {
	if !gm.noNeedToFree {
		C.Free(gm.m)
	}
}

func (gm GoMonoid) WalkChildren() MonoidResults {
	return *(*MonoidResults)(unsafe.Pointer(C.WalkChildren(gm.m)))
}

func (gm GoMonoid) WalkChildrenStack(results *MonoidResults) {
	C.WalkChildrenStack(gm.m, (*C.ulong)(unsafe.Pointer(&results[0])))
}

func (gm GoMonoid) RemoveGenerator(gen uint64) (res GoMonoid) {
	res.m = C.RemoveGenerator(gm.m, C.uint_fast64_t(gen))
	return
}

func (gm GoMonoid) Genus() uint64 {
	return uint64(C.Genus(gm.m))
}

func (gm GoMonoid) GetBytes() MonoidStorage {
	ms := make([]byte, C.SIZEOF_MONOID)
	copy(ms, (*[1 << 30]byte)(unsafe.Pointer(gm.m))[:C.SIZEOF_MONOID:C.SIZEOF_MONOID])
	return ms
}

func (gm GoMonoid) Print() {
	C.Print(gm.m)
}
