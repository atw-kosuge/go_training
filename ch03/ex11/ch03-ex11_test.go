// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"testing"
)

func Test_comma(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"", "1", "1"},
		{"", "123", "123"},
		{"", "1234", "1,234"},
		{"", "1234567", "1,234,567"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := comma(tt.s); got != tt.want {
				t.Errorf("comma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commaFloat(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"", "1", "1"},
		{"", "+1", "+1"},
		{"", "-1", "-1"},
		{"", "123", "123"},
		{"", "+123", "+123"},
		{"", "-123", "-123"},
		{"", "1234", "1,234"},
		{"", "+1234", "+1,234"},
		{"", "-1234", "-1,234"},
		{"", "1234567", "1,234,567"},
		{"", "+1234567", "+1,234,567"},
		{"", "-1234567", "-1,234,567"},
		{"", "1.1234567890", "1.1234567890"},
		{"", "+1.1234567890", "+1.1234567890"},
		{"", "-1.1234567890", "-1.1234567890"},
		{"", "123.1234567890", "123.1234567890"},
		{"", "+123.1234567890", "+123.1234567890"},
		{"", "-123.1234567890", "-123.1234567890"},
		{"", "1234.1234567890", "1,234.1234567890"},
		{"", "+1234.1234567890", "+1,234.1234567890"},
		{"", "-1234.1234567890", "-1,234.1234567890"},
		{"", "1234567.1234567890", "1,234,567.1234567890"},
		{"", "+1234567.1234567890", "+1,234,567.1234567890"},
		{"", "-1234567.1234567890", "-1,234,567.1234567890"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commaFloat(tt.s); got != tt.want {
				t.Errorf("commaFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
