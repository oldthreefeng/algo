/*
@Time : 2019/9/25 17:52
@Author : louis
@File : 02_test
@Software: GoLand
*/

package lintcode

import "testing"
var n int64 = 20

func TestTrailingZeros(t *testing.T) {
	//超过20, 21!溢出了,超过int64
	want := TrailingZerosX(n)
	rel  := TrailingZeros(n)
	if want != rel {
		t.Fatalf("want=%d, real=%d", want,rel)
	}
	t.Logf("want=%d", want)
}

func BenchmarkTrailingZeros(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TrailingZeros(n)
	}
}

func BenchmarkTrailingZerosX(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TrailingZerosX(n)
	}
}

/*
BenchmarkTrailingZeros-4    	12632070	        94.4 ns/op
BenchmarkTrailingZerosX-4   	209059014	         4.96 ns/op
*/