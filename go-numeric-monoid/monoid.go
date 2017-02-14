package nm

// #cgo LDFLAGS: -lcilkrts
// #cgo CPPFLAGS: -Wno-cpp -fPIC -g -O3 -fcilkplus -DNDEBUG -march=native -mtune=native
// #cgo CXXFLAGS: -std=c++11
// #include "ctreewalk.h"
import "C"

import (
	"runtime"
	"time"
	"unsafe"

	"github.com/intel-go/cpuid"
)

type GoMonoid struct {
	m C.Monoid
}

func init() {
	if C.MAX_GENUS < 16 {
		panic("This package has been compiled with MAX_GENUS<=16 but it is not supported")
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

func (gm GoMonoid) Walk() ([]uint, time.Duration) {
	start := time.Now()
	cres := C.WalkChildren(gm.m)
	// WARNING: casting C.ulong into uint is not guaranteed to be valid
	return (*[1 << 30]uint)(unsafe.Pointer(cres))[:C.MAX_GENUS:C.MAX_GENUS], time.Since(start)
}

func (gm GoMonoid) Free() {
	C.FreeMonoid(gm.m)
}
