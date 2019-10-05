/*
Copyright 2019 louis.
@Time : 2019/9/29 10:09
@Author : louis
@File : 66-preordertravelsal
@Software: GoLand

*/

package lintcode

import (
	"reflect"
	"testing"
)

func TestPreOrderTravelSal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test01", args{&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}}, []int{1, 2, 3}},
		{"test02", args{NewTravelSal()}, []int{1, 2, 3, 4, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PreOrderTravelSal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PreOrderTravelSal() = %v, want %v", got, tt.want)
			}
		})
	}
}
