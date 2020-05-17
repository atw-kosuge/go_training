package main

import (
	"fmt"
	"os"
	"sort"
)

// RuneSlice rune用のSortInterface実装
type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	if len(os.Args) > 2 {
		fmt.Printf("  %v\n", equalsForAnagram(os.Args[1], os.Args[2]))
	}
}

func equalsForAnagram(s1 string, s2 string) bool {
	runes1 := []rune(s1)
	runes2 := []rune(s2)

	// rune配列をソート
	sort.Sort(RuneSlice(runes1))
	sort.Sort(RuneSlice(runes2))

	for _, a := range runes1 {
		fmt.Print(a)
		fmt.Println(":" + string(a))
	}
	for _, a := range runes2 {
		fmt.Print(a)
		fmt.Println(":" + string(a))
	}

	// string変換してequals
	return string(runes1) == string(runes2)
}
