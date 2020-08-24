// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		// outline(url)

		targetID := "footer"
		elm, err := findById(url, targetID)
		if err == nil {
			fmt.Printf("%v is <%s>\n", targetID, elm.Data)
		}
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil && !pre(n) {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil && !post(n) {
		return
	}
}

func ElementByID(doc *html.Node, id string) (target *html.Node) {
	target = nil
	forEachNode(doc, func(doc *html.Node) bool {
		for _, attr := range doc.Attr {
			if attr.Key == "id" && attr.Val == id {
				target = doc
				return false
			}
		}
		return true //startElement(doc)
	}, nil)
	return
}

func findById(url string, id string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	elm := ElementByID(doc, id)

	return elm, nil
}
