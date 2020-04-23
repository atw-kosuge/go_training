package main

import (
	"fmt"
	"go_training/ch02/ex03/popcount"
	"time"
)

func main() {
	fmt.Printf("%d %dns\n", popcount.PopCount(138394), timeMeasurement(popcount.PopCount, 138394).Nanoseconds())
	fmt.Printf("%d %dns\n", popcount.LoopPopCount(138394), timeMeasurement(popcount.LoopPopCount, 138394).Nanoseconds())
}

func timeMeasurement(fn func(uint64) int, val uint64) time.Duration {
	start := time.Now()
	fn(val)
	return time.Since(start)
}
