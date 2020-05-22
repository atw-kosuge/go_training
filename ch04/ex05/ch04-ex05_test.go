package main

import (
	"testing"
)

func Test_removeSameChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"a"}, "a"},
		{"", args{"aa"}, "a"},
		{"", args{"aabbcc"}, "abc"},
		{"", args{"aaaabbbcc"}, "abc"},
		{"", args{"aaaabbbccd"}, "abcd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeSameChar(tt.args.s); got != tt.want {
				t.Errorf("removeSameChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeSameString(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{"", []string{"abcd"}, []string{"abcd"}},
		{"", []string{"abcd", "abcd"}, []string{"abcd"}},
		{"", []string{"abcd", "efgh", "abcd"}, []string{"abcd", "efgh", "abcd"}},
		{"", []string{"abcd", "abcd", "efgh", "abcd", "abcd"}, []string{"abcd", "efgh", "abcd"}},
		{"", []string{"abcd", "abcd", "efgh", "abcd", "abcd", "abcd"}, []string{"abcd", "efgh", "abcd"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeSameString(tt.args)
			if len(got) != len(tt.want) {
				t.Errorf("%v, %v", got, tt.want)
			}
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("%v, %v", got, tt.want)
				}
			}
		})
	}
}
