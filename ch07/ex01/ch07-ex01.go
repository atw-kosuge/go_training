package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// ByteCounter バイトカウンタ
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

// LineCounter 行カウンタ
type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += LineCounter(count)
	return count, scanner.Err()
}

// WordCounter Wordカウンタ
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += WordCounter(count)
	return count, scanner.Err()
}

func main() {
	var c ByteCounter
	t1 := "This is a pen"
	c.Write([]byte(t1))
	fmt.Printf("%v [byte: %v]\n", t1, c)

	var l LineCounter
	t2 := "This is\na\npen"
	l.Write([]byte(t2))
	fmt.Printf("%v [line: %v]\n", t2, l)

	var w WordCounter
	t3 := "This is a pen"
	w.Write([]byte(t3))
	fmt.Printf("%v [word: %v]\n", t3, w)
}
