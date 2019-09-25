/*
@Time : 2019/9/25 19:26
@Author : louis
@File : 04-uglynumber_test
@Software: GoLand
*/

package lintcode

import "testing"

func BenchmarkUglyNumber(b *testing.B) {
	n := 100
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UglyNumber(n)
	}
}

func TestUglyNumber(t *testing.T) {
	n, want := 9, 10
	rel := UglyNumber(n)
	if want != rel {
		t.Fatalf("want=%d, real=%d", want, rel)
	}
	t.Logf("want=%d", want)
}
