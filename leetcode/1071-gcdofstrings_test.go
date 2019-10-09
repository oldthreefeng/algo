/*
Copyright 2019 louis.
@Time : 2019/10/10 0:32
@Author : louis
@File : 1071-gcdofstrings
@Software: GoLand

*/

package leetcode

import "testing"

func TestGcdOfStrings(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test01",args{"ABCABC","ABC"},"ABC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GcdOfStrings(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GcdOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
