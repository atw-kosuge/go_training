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
