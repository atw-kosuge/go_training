package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()

	if len(flag.Args()) > 0 {
		fmt.Printf("%s\n", removeSameString(flag.Args()[0:]))
	} else {
		cmd := os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
		fmt.Fprintf(os.Stderr, "Usage %s: %s string\n", cmd, cmd)
		flag.PrintDefaults()
	}
}

func removeSameChar(s string) string {
	b := []byte(s)
	for i := 1; i < len(b); i++ {
		if b[i-1] == b[i] {
			copy(b[i-1:], b[i:])
			b = b[:len(b)-1]
			i--
		}
	}
	return string(b)
}

func removeSameString(s []string) []string {
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			copy(s[i-1:], s[i:])
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}
