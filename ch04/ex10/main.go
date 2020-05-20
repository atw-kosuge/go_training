// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"flag"
	"fmt"
	"go_training/ch04/ex10/github"
	"log"
	"os"
	"strings"
	"time"
)

//!+
func main() {
	flag.Parse()

	if len(flag.Args()) > 0 {
		printIssuesByPeriods(flag.Args())
	} else {
		cmd := os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
		fmt.Fprintf(os.Stderr, "Usage %s: %s string\n", cmd, cmd)
		flag.PrintDefaults()
	}
}

func printIssuesByPeriods(terms []string) {
	now := time.Now()

	p1 := now.AddDate(0, -1, 0)
	fmt.Printf("--- %s after ------------------ \n", p1.Format(time.RFC3339))
	printIssues(append(terms, fmt.Sprintf("created:>%s", p1.Format(time.RFC3339)), "sort:created-asc"))
	fmt.Println()

	p2 := now.AddDate(-1, 0, 0)
	fmt.Printf("--- %s after ------------------ \n", p2.Format(time.RFC3339))
	printIssues(append(terms, fmt.Sprintf("created:>%s", p2.Format(time.RFC3339)), "sort:created-asc"))
	fmt.Println()

	p3 := now.AddDate(-1, 0, 0)
	fmt.Printf("--- %s before ------------------ \n", p3.Format(time.RFC3339))
	printIssues(append(terms, fmt.Sprintf("created:<=%s", p3.Format(time.RFC3339)), "sort:created-asc"))
	fmt.Println()
}

func printIssues(terms []string) {
	fmt.Println(strings.Join(terms, " "))
	result, err := github.SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues: \n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("%s #%-5d %9.9s %.55s\n",
			item.CreatedAt.Format(time.RFC3339), item.Number, item.User.Login, item.Title)
	}
}
