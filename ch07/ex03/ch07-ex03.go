package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// ランダムデータ生成
	rand.Seed(time.Now().UnixNano())
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	fmt.Printf("origin: %v\n", strings.Trim(fmt.Sprint(data), "[]"))

	// tree
	var root *tree
	for _, v := range data {
		root = add(root, v)
	}
	fmt.Printf("sorted: %v\n", root.String())
}
