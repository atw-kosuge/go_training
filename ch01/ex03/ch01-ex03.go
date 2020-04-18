package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	inefficiencyTime := timeMeasurement(inefficiency).Nanoseconds()
	efficiencyTime := timeMeasurement(efficiency).Nanoseconds()

	fmt.Printf("%s: %d ns\n", "非効率", inefficiencyTime)
	fmt.Printf("%s: %d ns\n", "効率", efficiencyTime)
	fmt.Printf("%s: %d ns\n", "差分", inefficiencyTime-efficiencyTime)
}

func timeMeasurement(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}

func inefficiency() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func efficiency() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
