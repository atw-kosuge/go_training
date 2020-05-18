package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	use384 := flag.Bool("384", false, "use SHA384")
	use512 := flag.Bool("512", false, "use SHA512")
	flag.Parse()

	if len(flag.Args()) > 0 {
		if *use384 {
			fmt.Printf("SHA384 %x\n", sha512.Sum384([]byte(flag.Arg(0))))
		} else if *use512 {
			fmt.Printf("SHA512 %x\n", sha512.Sum512([]byte(flag.Arg(0))))
		} else {
			fmt.Printf("SHA256 %x\n", sha256.Sum256([]byte(flag.Arg(0))))
		}
	} else {
		cmd := os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
		fmt.Fprintf(os.Stderr, "Usage %s: %s string\n", cmd, cmd)
		flag.PrintDefaults()
	}
}
