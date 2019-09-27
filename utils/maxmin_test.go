/*
@Time : 2019/9/19 0:34
@Author : louis
@File : maxmin
@Software: GoLand
*/

package utils

import "testing"

func TestPower(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01",args{2,2},4},
		{"test02",args{3,2},9},
		{"test03",args{3,0},1},
		{"test04",args{33,3},35937},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Power(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Power() = %v, want %v", got, tt.want)
			}
		})
	}
}
