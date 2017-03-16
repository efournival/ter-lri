package nm

// #cgo LDFLAGS: -lcilkrts
// #cgo CXXFLAGS: -std=c++11 -Wno-cpp -fPIC -g -O3 -fcilkplus -march=native -mtune=native
// #include "ctreewalk.h"
import "C"

import (
	"runtime"
	"unsafe"

	"github.com/intel-go/cpuid"
)

type (
	MonoidResults [C.MAX_GENUS]uint64

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

	// This is not mandatory but it helps with timing consistency
	runtime.LockOSThread()
}

func NewMonoid() GoMonoid {
	var gm GoMonoid
	gm.m = C.InitFullN()
	return gm
}

func (gm GoMonoid) Walk() []uint64 {
	cres := C.WalkChildren(gm.m)
	return (*[1 << 30]uint64)(unsafe.Pointer(cres))[:C.MAX_GENUS:C.MAX_GENUS]
}

func (gm GoMonoid) WalkChildrenStack(results *MonoidResults) {
	C.WalkChildrenStack(gm.m, (*C.ulong)(unsafe.Pointer(&results[0])))
}

func (gm GoMonoid) RemoveGenerator(gen uint) (res GoMonoid) {
	res.m = C.RemoveGenerator(gm.m, C.uint(gen))
	return
}

func (gm GoMonoid) Genus() uint64 {
	return uint64(C.Genus(gm.m))
}

func (gm GoMonoid) Free() {
	C.FreeMonoid(gm.m)
}
