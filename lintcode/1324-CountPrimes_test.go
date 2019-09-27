/*
Copyright 2019 louis.
@Time : 2019/9/26 23:18
@Author : louis
@File : 1324-CountPrimes_test
@Software: GoLand

*/

package lintcode

import "testing"

/*
计算小于非负数n的质数的个数。

Example
样例 1

输入: n = 2
输出: 0
样例 2

输入: n = 4
输出: 2
解析：2, 3 是素数
*/

func TestCountPrimes(t *testing.T) {
	testClass := []struct {
		n     int
		count int
	}{
		{2, 0},
		{4, 2},
		{100, 25},
		{1000, 168},
		{10000, 1229},
	}
	for _,tc :=range  testClass {
		rel :=CountPrimes(tc.n)
		want := tc.count
		if want != rel {
			t.Fatalf("want=%d, real=%d", want,rel)
		}
		t.Logf("want=%d", want)
	}
}

func BenchmarkCountPrimes(b *testing.B) {
	var n = 10000
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountPrimes(n)
	}
}