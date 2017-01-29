package main

import (
	"fmt"
	"runtime"
)

const MAX_GENUS = 29

func main() {
	fmt.Printf("Maximum genus: %d\nMaximum procs to be spawned: %d\n", MAX_GENUS, runtime.GOMAXPROCS(0))
	NewPool().Start()
}
