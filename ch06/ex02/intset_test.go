// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"reflect"
	"testing"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestIntSet_Len(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{"no value", []int{}, 0},
		{"one value", []int{1}, 1},
		{"two value", []int{1, 2}, 2},
		{"境界値", []int{1, 64}, 2},
		{"境界値＋１", []int{1, 65}, 2},
		{"境界値＋１　複数値", []int{1, 65, 70}, 3},
		{"複数値", []int{1, 65, 129, 192, 193}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s IntSet
			for _, value := range tt.values {
				s.Add(value)
			}
			if got := s.Len(); got != tt.want {
				t.Errorf("IntSet.Len(%v) = %v, want %v", tt.values, got, tt.want)
			}
		})
	}
}

func TestIntSet_Remove(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		arg    int
		want   string
	}{
		{"値なしから削除", []int{}, 1, "{}"},
		{"対象値削除", []int{1}, 1, "{}"},
		{"複数値状態から削除", []int{1, 2}, 2, "{1}"},
		{"境界値削除", []int{1, 5, 64}, 64, "{1 5}"},
		{"境界値+1削除", []int{1, 5, 64, 65}, 65, "{1 5 64}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s IntSet
			for _, value := range tt.values {
				s.Add(value)
			}
			s.Remove(tt.arg)
			if got := s.String(); got != tt.want {
				t.Errorf("IntSet(%v).Remove(%v) = %v, want %v", tt.values, tt.arg, got, tt.want)
			}
		})
	}
}

func TestIntSet_Clear(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   string
	}{
		{"値なしからクリア", []int{}, "{}"},
		{"クリア1", []int{1}, "{}"},
		{"クリア2", []int{1, 64}, "{}"},
		{"クリア3", []int{1, 64, 65}, "{}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s IntSet
			for _, value := range tt.values {
				s.Add(value)
			}
			s.Clear()
			if got := s.String(); got != tt.want {
				t.Errorf("IntSet(%v).Clear() = %v, want %v", tt.values, got, tt.want)
			}
		})
	}
}

func TestIntSet_Copy(t *testing.T) {
	tests := []struct {
		name   string
		values []int
	}{
		{"no value", []int{}},
		{"one value", []int{1}},
		{"two value", []int{1, 2}},
		{"境界値", []int{1, 64}},
		{"境界値＋１", []int{1, 65}},
		{"境界値＋１　複数値", []int{1, 65, 70}},
		{"複数値", []int{1, 65, 129, 192, 193}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s IntSet
			for _, value := range tt.values {
				s.Add(value)
			}
			if got := s.Copy(); got.String() != s.String() || reflect.DeepEqual(got, s) {
				t.Errorf("IntSet.Copy() = %v, want %v", got.String(), s.String())
			}
		})
	}
}
