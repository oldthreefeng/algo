/*
Copyright 2019 louis.
@Time : 2019/10/7 23:51
@Author : louis
@File : 476-numbercomplement
@Software: GoLand

*/

package leetcode

import "testing"

func TestFindComplement(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01",args{5},2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindComplement(tt.args.num); got != tt.want {
				t.Errorf("FindComplement() = %v, want %v", got, tt.want)
			}
		})
	}
}
