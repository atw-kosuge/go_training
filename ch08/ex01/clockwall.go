// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 221.
//!+

// Netcat1 is a read-only TCP client.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"regexp"
)

type clockClient struct {
	name string
	url  string
}

func main() {
	flag.Parse()
	clients := parseArgs(flag.Args())
	done := make(chan string)
	for idx, client := range clients {
		go netcat(idx, client, done)
	}
	for {
		test := <-done
		fmt.Print(test)

	}
	// <-done
}

func parseArgs(args []string) []clockClient {
	var clients []clockClient
	for _, arg := range args {
		pattern := regexp.MustCompile(`(\S+)=(\S+)`)
		submatch := pattern.FindStringSubmatch(arg)
		clients = append(clients, clockClient{submatch[1], submatch[2]})
	}
	return clients
}

func netcat(idx int, client clockClient, out chan<- string) {
	conn, err := net.Dial("tcp", client.url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		out <- fmt.Sprintf("%s: %s\n", client.name, string(bytes))
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

//!-
