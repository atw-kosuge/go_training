package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	printText(doc)
}

func printText(n *html.Node) {
	if n != nil {
		if n.Type == html.TextNode {
			//fmt.Println(n.Data)
			if n.Parent.Data != "script" && n.Parent.Data != "style" {
				text := strings.TrimSpace(n.Data)
				if len(text) > 0 {
					fmt.Println(text)
				}
			}
		} else {
			printText(n.FirstChild)
		}
		printText(n.NextSibling)
	}
}
