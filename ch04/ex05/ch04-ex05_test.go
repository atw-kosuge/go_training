package main

import "testing"

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
