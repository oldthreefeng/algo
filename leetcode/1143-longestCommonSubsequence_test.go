/*
@Time : 2019/10/8 14:32
@Author : louis
@File : 1143-longestCommonSubsequence
@Software: GoLand
*/

package leetcode

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01",args{"abcde","ace"},3},
		{"test02",args{"abc","def"},0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lcs(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("Lcs() = %v, want %v", got, tt.want)
			}
		})
	}
}
