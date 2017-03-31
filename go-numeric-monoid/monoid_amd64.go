package nm

/*
	#cgo LDFLAGS: -lcilkrts
	#cgo CXXFLAGS: -std=c++11 -Wno-cpp -fPIC -g -O3 -fcilkplus -march=native -mtune=native
	#include "ctreewalk.h"
	#include <stdlib.h>
*/
import "C"

import (
	"runtime"
	"unsafe"

	"github.com/intel-go/cpuid"
)

type (
	MonoidResults [C.MAX_GENUS]uint64

	// Raw array size because of:
	// https://github.com/golang/go/issues/19816
	MonoidStorage [144]byte

	GoMonoid struct {
		m C.Monoid
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

	runtime.SetFinalizer(&gm, func(m *GoMonoid) {
		C.free(unsafe.Pointer(gm.m))
	})

	return gm
}

func FromBytes(ms MonoidStorage) GoMonoid {
	var gm GoMonoid
	gm.m = C.Monoid(&ms[0])
	return gm
}

func (gm GoMonoid) Walk() []uint64 {
	return (*[1 << 30]uint64)(unsafe.Pointer(C.WalkChildren(gm.m)))[:C.MAX_GENUS:C.MAX_GENUS]
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

func (gm GoMonoid) GetBytes() (ms MonoidStorage) {
	copy((*MonoidStorage)(unsafe.Pointer(gm.m))[0:C.SIZEOF_MONOID], ms[:])
	return
}
