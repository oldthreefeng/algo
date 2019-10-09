/*
Copyright 2019 louis.
@Time : 2019/10/9 22:54
@Author : louis
@File : 5-longestPalindrome
@Software: GoLand

*/

package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test01",args{"cbbd"},"bb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("LongestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
