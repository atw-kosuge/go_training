// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", commaFloat(os.Args[i]))
	}
}

func commaFloat(s string) string {
	dot := strings.Index(s, ".")
	if dot == -1 {
		return comma(s)
	}
	regular := s[:dot]
	decimal := s[dot:]
	return comma(regular) + decimal
}

func comma(s string) string {
	var buf bytes.Buffer
	if !('0' <= s[0] && s[0] <= '9') {
		buf.WriteString(string(s[0]))
		s = s[1:]
	}

	n := len(s)
	if n <= 3 {
		buf.WriteString(s)
	} else {
		surplus := n % 3
		if surplus > 0 {
			buf.WriteString(s[:surplus])
		}

		split := n / 3
		for i := 0; i < split; i++ {
			if (i == 0 && surplus > 0) || i > 0 {
				buf.WriteString(",")
			}
			buf.WriteString(s[3*i+surplus : 3*i+3+surplus])
		}
	}
	return buf.String()
}
