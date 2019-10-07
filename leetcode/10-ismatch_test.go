/*
Copyright 2019 louis.
@Time : 2019/10/7 15:09
@Author : louis
@File : 10-ismatch
@Software: GoLand

*/

package leetcode

import "testing"

func TestIsMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test01",args{"aa","a*"},true},
		{"test02",args{"aa","a"},false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("IsMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
