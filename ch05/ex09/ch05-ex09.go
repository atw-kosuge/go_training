package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	for _, str := range os.Args[1:] {
		replaced := expand(str, func(target string) string {
			return "*"
		})
		fmt.Println(replaced)
	}
}

func expand(s string, f func(string) string) string {
	pat := regexp.MustCompile(`(\$\S+)`)
	return pat.ReplaceAllStringFunc(s, func(m string) string {
		submatch := pat.FindStringSubmatch(m)
		return f(submatch[1][1:])
	})
}
