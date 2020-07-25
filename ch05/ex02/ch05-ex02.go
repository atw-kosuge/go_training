package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for key, count := range countTag(doc, nil) {
		fmt.Printf("Tag[%s]:%d\n", key, count)
	}
}

func countTag(n *html.Node, result map[string]int) map[string]int {
	if result == nil {
		result = make(map[string]int)
	}
	if n != nil {
		if n.Type == html.ElementNode {
			result[n.Data]++
		}
		countTag(n.FirstChild, result)
		countTag(n.NextSibling, result)
	}
	return result
}
