/*
@Time : 2019/10/9 11:52
@Author : louis
@File : 680-validpalindrome2
@Software: GoLand
*/

package leetcode

import "testing"

func TestValidPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test01",args{"abca"},true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidPalindrome(tt.args.s); got != tt.want {
				t.Errorf("ValidPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
