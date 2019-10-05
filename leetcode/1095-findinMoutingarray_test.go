/*
Copyright 2019 louis.
@Time : 2019/9/30 20:36
@Author : louis
@File : 1095-findinMountingarray
@Software: GoLand

*/

package leetcode

import "testing"

func TestFindInMountainArray(t *testing.T) {
	type args struct {
		target int
		m      *MountainArray
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01",args{3,&MountainArray{[]int{1,2,3,4,5,3,1}}},2},
		{"test02",args{3,&MountainArray{[]int{1,2,5,7,9,2,1}}},-1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindInMountainArray(tt.args.target, tt.args.m); got != tt.want {
				t.Errorf("FindInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
