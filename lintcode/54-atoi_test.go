/*
Copyright 2019 louis.
@Time : 2019/9/28 17:54
@Author : louis
@File : 54-atoi
@Software: GoLand

*/

package lintcode

import "testing"

func TestAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01",args{"+1"},1},
		{"test02",args{"10"},10},
		{"test03",args{"     010"},10},
		{"test04",args{" +123k "},123},
		{"test05",args{"1.0"},1},
		{"test06",args{" 15+4"},15},
		{"test07",args{"   +-1111 "},0},
		{"test08",args{"1234567890123456789012345678901234567890"},INT_MAX},
		{"test09",args{"-123123123123123"},INT_MIN},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Atoi(tt.args.s); got != tt.want {
				t.Errorf("Atoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
