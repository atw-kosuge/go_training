// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go_training/ch04/ex11/github"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()

	if len(flag.Args()) > 3 {
		v3client := github.V3Client{Token: flag.Arg(1), Owner: flag.Arg(2), Repository: flag.Arg(3)}
		switch flag.Arg(0) {
		case "create":
			createIssue(v3client, flag.Args()[4:])
		case "edit":
			editIssue(v3client, flag.Args()[4:])
		case "close":
			closeIssue(v3client, flag.Args()[4:])
		default:
			printUsage()
			flag.PrintDefaults()
		}

		//client := authenticate("8f3ddf2ff691da2af2032709fc16dc0222ddd5d5")
		//v3client.Authenticate()

		// {
		// 	issue, err := v3client.CreateIssue("test", "bodytest")
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}

		// 	data, err := json.MarshalIndent(*issue, "", "    ")
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Printf("%s\n", data)
		// }
		// {
		// 	issue, err := v3client.EditIssue(3, "test", "bodytest-edit")
		// 	//issue, err := v3client.CloseIssue(3)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}

		// 	data, err := json.MarshalIndent(*issue, "", "    ")
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Printf("%s\n", data)
		// }

		//createIssue("8f3ddf2ff691da2af2032709fc16dc0222ddd5d5", "atware-a-kosuge", "go_training", "test", "bodytest")
	} else {
		printUsage()
		flag.PrintDefaults()
	}
}

func printUsageCreate() {
	cmd := os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
	fmt.Fprintf(os.Stderr, "Usage %s create:\n", cmd)
	fmt.Fprintf(os.Stderr, "    %s create token owner repository title body\n", cmd)
}

func printUsageEdit() {
	cmd := os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
	fmt.Fprintf(os.Stderr, "Usage %s edit:\n", cmd)
	fmt.Fprintf(os.Stderr, "    %s edit token owner repository number title body\n", cmd)
}

func printUsageClose() {
	cmd := os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
	fmt.Fprintf(os.Stderr, "Usage %s close:\n", cmd)
	fmt.Fprintf(os.Stderr, "    %s close token owner repository number\n", cmd)
}

func printUsage() {
	cmd := os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
	fmt.Fprintf(os.Stderr, "Usage %s:\n", cmd)
	fmt.Fprintf(os.Stderr, "    %s create|edit|close\n", cmd)
}

func createIssue(v3client github.V3Client, args []string) {
	if len(args) > 1 {
		issue, err := v3client.CreateIssue(args[0], args[1])
		if err != nil {
			log.Fatal(err)
			return
		}

		data, err := json.MarshalIndent(*issue, "", "    ")
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("%s\n", data)
	} else {
		printUsageCreate()
	}
}

func editIssue(v3client github.V3Client, args []string) {
	if len(args) > 2 {
		number, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
			printUsageEdit()
		} else {
			issue, err := v3client.EditIssue(number, args[1], args[2])
			if err != nil {
				log.Fatal(err)
				return
			}

			data, err := json.MarshalIndent(*issue, "", "    ")
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Printf("%s\n", data)
		}
	} else {
		printUsageEdit()
	}
}

func closeIssue(v3client github.V3Client, args []string) {
	if len(args) > 0 {
		number, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
			printUsageClose()
		} else {
			issue, err := v3client.CloseIssue(number)
			if err != nil {
				log.Fatal(err)
				return
			}

			data, err := json.MarshalIndent(*issue, "", "    ")
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Printf("%s\n", data)
		}
	} else {
		printUsageClose()
	}
}
