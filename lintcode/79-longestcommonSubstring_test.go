/*
Copyright 2019 louis.
@Time : 2019/10/9 23:45
@Author : louis
@File : 79-longestcommonSubstring
@Software: GoLand

*/

package lintcode

import "testing"

func TestLss(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01",args{"abcd","cbce"},2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lss(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Lss() = %v, want %v", got, tt.want)
			}
		})
	}
}
