package main

import "testing"

func Test_compressSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"a"}, "a"},
		{"", args{"a　a"}, "a　a"},
		{"", args{"　a　"}, "　a　"},
		{"", args{"a　　b"}, "a b"},
		{"", args{"a　　　b　　　　c"}, "a b c"},
		{"", args{"a　  　　b 　　 　　 c"}, "a b c"},
		{"", args{"  a　  　　b 　　 　　 c   "}, " a b c "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressSpace(tt.args.s); got != tt.want {
				t.Errorf("removeUnicodeSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
