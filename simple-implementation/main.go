package main

import "fmt"

const (
	MAX_GENUS  = 10
	COUNT_SIZE = MAX_GENUS + 1
	SIZE       = 3 * MAX_GENUS
)

func main() {
	fmt.Printf("Naive:\n  %+v\n", Count(InitNaiveMonoid()))
	fmt.Printf("Optimized:\n  %+v\n", Count(InitOptimizedMonoid()))
}
