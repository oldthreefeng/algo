/*
Copyright 2019 louis.
@Time : 2019/9/28 22:01
@Author : louis
@File : 56-Twosum
@Software: GoLand

*/

package lintcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		a []int
		m int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test02", args{[]int{2, 7, 9, 12}, 9}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.a, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoSumSlow(t *testing.T) {
	type args struct {
		a []int
		m int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test02", args{[]int{1, 0, -1}, 0}, []int{0, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSumSlow(tt.args.a, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSumSlow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	a := []int{2, 3, 9, 12, 15, 10, 5, 1, 7}
	m := 9
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum(a, m)
	}
}

func BenchmarkTwoSumSlow(b *testing.B) {
	a := []int{2, 3, 9, 12, 15, 10, 5, 1, 7}
	m := 9
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSumSlow(a, m)
	}
}

/*
goos: windows
goarch: amd64
pkg: github.com/oldthreefeng/algo/lintcode
BenchmarkTwoSum-4       	 3438445	       327 ns/op
BenchmarkTwoSumSlow-4   	27269504	        50.1 ns/op
*/