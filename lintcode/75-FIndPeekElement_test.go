/*
Copyright 2019 louis.
@Time : 2019/9/29 15:51
@Author : louis
@File : 75-FIndPeekElement
@Software: GoLand

*/

package lintcode

import "testing"

func TestFindPeak(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01",args{[]int{1,2,1,3,4,6,7,6}},1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPeak(tt.args.A); got != tt.want {
				t.Errorf("FindPeak() = %v, want %v", got, tt.want)
			}
		})
	}
}
