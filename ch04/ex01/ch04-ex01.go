package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Printf("Count=%d\n", countBitDiffBySha256(os.Args[1], os.Args[2]))
	}
}

func countBitDiffBySha256(hash1, hash2 string) int {
	c1 := sha256.Sum256([]byte(hash1))
	c2 := sha256.Sum256([]byte(hash2))
	count := 0
	for i := range c1 {
		xor := c1[i] ^ c2[i]
		bitcnt := bitCount(xor)
		count += bitcnt

		fmt.Printf("hash1[%d]: %08b\n", i, c1[i])
		fmt.Printf("hash2[%d]: %08b\n", i, c2[i])
		fmt.Printf("      xor: %08b count=%d total=%d\n", xor, bitcnt, count)
	}
	return count
}

func bitCount(x byte) int {
	count := 0
	for i := 0; i < 8; i++ {
		count += int(byte(x >> i & 1))
	}
	fmt.Printf("%b:%d\n", x, count)
	return count
}
