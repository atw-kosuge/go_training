package main

import "testing"

func Test_equalsAnagram(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"a", "a"}, true},
		{"", args{"asdf", "fdsa"}, true},
		{"", args{"asdf", "sfad"}, true},
		{"", args{"asdf", "asdfg"}, false},
		{"", args{"あ", "あ"}, true},
		{"", args{"世界", "界世"}, true},
		{"", args{"Hello,世界", "o,l世Hl界e"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalsForAnagram(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("equalsAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
