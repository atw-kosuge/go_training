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
		countMap := countUnicode(flag.Arg(0))
		for index, value := range countMap {
			fmt.Printf("[%s] = %d\n", index, value)
		}
	} else {
		cmd := os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
		fmt.Fprintf(os.Stderr, "Usage %s: %s string\n", cmd, cmd)
		flag.PrintDefaults()
	}
}

func countUnicode(s string) map[string]int {
	result := make(map[string]int)

	runes := []rune(s)
	for _, r := range runes {
		if unicode.IsDigit(r) {
			result["Digit"]++
		}
		if unicode.IsLetter(r) {
			result["Leter"]++
		}
		if unicode.IsMark(r) {
			result["Mark"]++
		}
	}
	return result
}
