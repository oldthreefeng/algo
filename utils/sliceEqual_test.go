/*
@Time : 2019/9/26 14:08
@Author : louis
@File : sliceEqual_test
@Software: GoLand
*/

package utils


import "testing"

func BenchmarkStringSliceEqual(b *testing.B) {
	sa := []string{"q", "w", "e", "r", "t"}
	sb := []string{"q", "w", "e", "r", "t", "x"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		StringSliceEqual(sa, sb)
	}
}

func BenchmarkStringSliceReflectEqual(b *testing.B) {
	sa := []string{"q", "w", "e", "r", "t"}
	sb := []string{"q", "w", "e", "r", "t", "x"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		StringSliceReflectEqual(sa, sb)
	}
}

func BenchmarkStringSliceEqualBce(b *testing.B) {
	sa := []string{"q", "w", "e", "r", "t"}
	sb := []string{"q", "w", "e", "r", "t", "x"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		StringSliceEqualBce(sa, sb)
	}
}

/*
BenchmarkStringSliceEqual-4          	252633386	         4.56 ns/op
BenchmarkStringSliceReflectEqual-4   	 9022603	       161 ns/op
BenchmarkStringSliceEqualBce-4       	385856928	         3.14 ns/op
*/

