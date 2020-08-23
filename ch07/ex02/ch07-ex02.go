package main

import (
	"fmt"
	"os"
)

func main() {
	w, cnt := CountingWriter(os.Stdout)

	t1 := "test abce"
	n, err := w.Write([]byte(t1))
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf(" [count:%v total:%v]\n", n, *cnt)

	t2 := "1234567890"
	n, err = w.Write([]byte(t2))
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf(" [count:%v total:%v]\n", n, *cnt)

	t3 := "12345"
	n, err = w.Write([]byte(t3))
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf(" [count:%v total:%v]\n", n, *cnt)
}
