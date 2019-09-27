/*
Copyright 2019 louis.
@Time : 2019/9/27 23:36
@Author : louis
@File : 135-combinationsum
@Software: GoLand

*/

package lintcode

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{

		{"test01", args{[]int{}, 3}, [][]int{}},
		{"test02", args{[]int{3, 12, 9, 11, 6, 7, 8, 5, 4}, 15}, [][]int{
			{3, 3, 3, 3, 3}, {3, 3, 3, 6}, {3, 3, 4, 5}, {3, 3, 9}, {3, 4, 4, 4},
			{3, 4, 8}, {3, 5, 7}, {3, 6, 6}, {3, 12}, {4, 4, 7}, {4, 5, 6}, {4, 11},
			{5, 5, 5}, {6, 9}, {7, 8},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CombinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCombinationSumX(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test01", args{[]int{}, 1}, [][]int{}},
		{"test02", args{[]int{7,1,2,5,1,6,10}, 8}, [][]int{
			{1,1,6},{1,2,5},{1,7},{2,6},
		}},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombinationSumX(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CombinationSumX() = %v, want %v", got, tt.want)
			}
		})
	}
}
