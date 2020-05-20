package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	flag.Parse()

	if len(flag.Args()) > 0 {
		fmt.Printf("%s\n", reverse(flag.Arg(0)))
	} else {
		cmd := os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
		fmt.Fprintf(os.Stderr, "Usage %s: %s string\n", cmd, cmd)
		flag.PrintDefaults()
	}
}

func reverse(s string) string {
	b := []byte(s)
	reverseUTF8bytes(b)
	return string(b)
}

func reverseUTF8bytes(b []byte) {
	size := len(b)
	n := 0
	for n < size {
		tmp, rlen := utf8.DecodeLastRune(b)
		fmt.Printf("%v %v %v\n", string(b), rlen, n)
		copy(b[rlen+n:], b[n:])
		copy(b[n:], []byte(string(tmp)))
		n += rlen
	}
}
