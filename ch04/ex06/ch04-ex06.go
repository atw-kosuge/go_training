package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	flag.Parse()

	if len(flag.Args()) > 0 {
		fmt.Printf("%s\n", removeUnicodeSpace(flag.Arg(0)))
	} else {
		cmd := os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
		fmt.Fprintf(os.Stderr, "Usage %s: %s string\n", cmd, cmd)
		flag.PrintDefaults()
	}
}

func removeUnicodeSpace(s string) string {
	r := []rune(s)
	for i := 1; i < len(r); i++ {
		if unicode.IsSpace(r[i-1]) && unicode.IsSpace(r[i]) {
			r[i] = rune(0x20)
			copy(r[i-1:], r[i:])
			r = r[:len(r)-1]
			i--
		}
		fmt.Println(string(r))
	}
	return string(r)
}
