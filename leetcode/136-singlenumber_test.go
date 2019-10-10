/*
@Time : 2019/10/9 16:52
@Author : louis
@File : 136-singlenumber_test.go
@Software: GoLand
*/

package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func BenchmarkSingleNumber2(b *testing.B) {
	var a = []int{1, 1, 1, 3, 3, 3, 5, 5, 5, 2, 6, 6, 6, 8, 8, 8}
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		SingleNumber2(a)
	}
}

func BenchmarkSingleNumber2x(b *testing.B) {
	var a = []int{1, 1, 1, 3, 3, 3, 5, 5, 5, 2, 6, 6, 6, 8, 8, 8}
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		SingleNumber2x(a)
	}
}

func TestSingleNumber3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test01",args{[]int{1,2,1,3,2,5}},[]int{3,5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SingleNumber3(tt.args.nums)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SingleNumber3() = %v, want %v", got, tt.want)
			}
		})
	}
}
