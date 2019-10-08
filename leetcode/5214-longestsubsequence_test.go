/*
@Time : 2019/10/8 13:57
@Author : louis
@File : 5214-longestsubsequence
@Software: GoLand
*/

package leetcode

import "testing"

func TestLongestSubsequence(t *testing.T) {
	type args struct {
		arr        []int
		difference int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01",args{[]int{1,2,3,4},1},4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestSubsequence(tt.args.arr, tt.args.difference); got != tt.want {
				t.Errorf("LongestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
