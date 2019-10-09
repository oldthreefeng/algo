/*
@Time : 2019/10/9 10:49
@Author : louis
@File : 1216-isValidPalindrome
@Software: GoLand
*/

package leetcode

import "testing"

func TestIsValidPalindrome(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test01",args{"abcdeca",2},true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidPalindrome(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("IsValidPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
