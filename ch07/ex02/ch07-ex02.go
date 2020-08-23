// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"fmt"
	"os"
)

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
