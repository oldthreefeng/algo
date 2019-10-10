/*
@Time : 2019/10/10 17:07
@Author : louis
@File : 20191010-1
@Software: GoLand
*/

package leetcode

import (
	"reflect"
	"testing"
)

func TestFraction(t *testing.T) {
	type args struct {
		cont []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test01", args{[]int{3, 2, 0, 2}}, []int{13, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fraction(tt.args.cont); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fraction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFractionX(t *testing.T) {
	type args struct {
		cont []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test01", args{[]int{3, 2, 0, 2}}, []int{13, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FractionX(tt.args.cont); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FractionX() = %v, want %v", got, tt.want)
			}
		})
	}
}
