/*
Copyright 2019 louis.
@Time : 2019/10/1 21:36
@Author : louis
@File : 941-vaildmoutingarray
@Software: GoLand

*/

package leetcode

import "testing"

func TestValidMountingArray(t *testing.T) {
	type args struct {
		m []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test01",args{[]int{0,1,2,3,1}},true},
		{"test02",args{[]int{2,1}},false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidMountingArray(tt.args.m); got != tt.want {
				t.Errorf("ValidMountingArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
