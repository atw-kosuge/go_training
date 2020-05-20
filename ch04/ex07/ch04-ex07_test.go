package main

import (
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"a"}, "a"},
		{"", args{"abcdef"}, "fedcba"},
		{"", args{"あ"}, "あ"},
		{"", args{"あいうえお"}, "おえういあ"},
		{"", args{"これはabcあれは何だ"}, "だ何はれあcbaはれこ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.s); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
