/*
@Time : 2019/10/9 12:13
@Author : louis
@File : 392-isSubsequence
@Software: GoLand
*/

package leetcode

import "testing"

func TestIsSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test01",args{"abc","achbekjc"},true},
		{"test02",args{"abc","ahbgdc"},true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
