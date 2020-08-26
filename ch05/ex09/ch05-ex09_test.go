package main

import (
	"strings"
	"testing"
)

func Test_expand(t *testing.T) {
	type args struct {
		s string
		f func(string) string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"$test", strings.ToUpper}, "TEST"},
		{"2", args{"aa$test", strings.ToUpper}, "aaTEST"},
		{"3", args{"aa$test aa", strings.ToUpper}, "aaTEST aa"},
		{"4", args{"$test $abcd", strings.ToUpper}, "TEST ABCD"},
		{"5", args{"", strings.ToUpper}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expand(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("expand() = %v, want %v", got, tt.want)
			}
		})
	}
}
