/*
@Time : 2019/9/25 17:36
@Author : louis
@File : lin_test
@Software: GoLand
*/

package lintcode

import "testing"

func TestDigitCounts(t *testing.T) {
	k, n, want := 1, 12, 5

	if DigitCounts(k, n) != want {
		t.Fatalf("want=%d, real=%d", want, DigitCounts(k, n))
	}

}

func BenchmarkDigitCounts(b *testing.B) {
	k, n := 1, 12
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		DigitCounts(k, n)
	}

}
