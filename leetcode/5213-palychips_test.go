/*
Copyright 2019 louis.
@Time : 2019/10/7 23:34
@Author : louis
@File : 5213-palychips
@Software: GoLand

*/

package leetcode

import "testing"

func TestMinCostToMoveChips(t *testing.T) {
	type args struct {
		chips []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01",args{[]int{1,2,3}},1},
		{"test02",args{[]int{2,2,2,3,3}},2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinCostToMoveChips(tt.args.chips); got != tt.want {
				t.Errorf("MinCostToMoveChips() = %v, want %v", got, tt.want)
			}
		})
	}
}
