/*
@Time : 2019/9/26 14:03
@Author : louis
@File : 06-mergeSortedArray_test
@Software: GoLand
*/

package lintcode

import (
	"github.com/oldthreefeng/algo/utils"
	"testing"
)

var (
	a1 = []int{1, 2, 3, 4}
	a2 = []int{2, 4, 5, 6}
)

func TestMergeSortedArray(t *testing.T) {

	want := []int{1, 2, 2, 3, 4, 4, 5, 6}
	rel := MergeSortedArray(a1, a2)
	if !utils.IntSliceEqualBce(want, rel) {
		t.Fatalf("want=%d, real=%d", want, rel)
	}
	t.Logf("rel=%d", rel)
}

func BenchmarkMergeSortedArray(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSortedArray(a1, a2)
	}
}

func BenchmarkMergeArrayX(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeArrayX(a1, a2)
	}
}

/*
BenchmarkMergeSortedArray-4   	16828744	        74.4 ns/op
BenchmarkMergeArrayX-4        	 3157518	       405 ns/op
*/