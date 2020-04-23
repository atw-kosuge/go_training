package main

import (
	"fmt"
	"go_training/ch02/ex01/tempconv"
	"reflect"
	"runtime"
)

func main() {
	fmt.Printf("%s: %f\n", getFunctionName(tempconv.CToF), tempconv.CToF(1.0))
	fmt.Printf("%s: %f\n", getFunctionName(tempconv.FToC), tempconv.CToF(1.0))
	fmt.Printf("%s: %f\n", getFunctionName(tempconv.CToK), tempconv.CToK(1.0))
	fmt.Printf("%s: %f\n", getFunctionName(tempconv.FToK), tempconv.FToK(1.0))
	fmt.Printf("%s: %f\n", getFunctionName(tempconv.KToC), tempconv.KToC(1.0))
	fmt.Printf("%s: %f\n", getFunctionName(tempconv.KToF), tempconv.KToF(1.0))
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
