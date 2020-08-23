// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"fmt"
	"os"
)

// type CountingWriter io.Writer

// var writer io.Writer
// var counter int64

// type CountingWriter struct {
// 	writer       io.Writer
// 	bytesWritten int
// }

// func (c CountingWriter) Write(p []byte) (int, error) {
// 	n, err := writer.Write(p)
// 	counter += int64(n)
// 	return n, err
// }

// func CountingWriter(w io.Writer) (io.Writer, *int64) {

// 	return CountingWriter, &counter
// }

func main() {

	w, cnt := CountingWriter(os.Stdout)

	n, err := w.Write([]byte("test\n"))
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("write:%v total:%v %v\n", n, *cnt, cnt)

	n, err = w.Write([]byte("test2\n"))
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("write:%v total:%v %v\n", n, *cnt, cnt)
}
