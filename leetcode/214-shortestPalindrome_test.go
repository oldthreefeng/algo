/*
Copyright 2019 louis.
@Time : 2019/10/9 22:06
@Author : louis
@File : 214-shortestPalindrome_test.go
@Software: GoLand

*/

package leetcode

import "testing"

func BenchmarkShortestPalindrome(b *testing.B) {
	s := "aacecaaa"
	b.ResetTimer()
	for i:=0;i<b.N;i++ {
		ShortestPalindrome(s)
	}
}

func BenchmarkShortestPalindromeX(b *testing.B) {
	s := "aacecaaa"
	b.ResetTimer()
	for i:=0;i<b.N;i++ {
		ShortestPalindromeX(s)
	}
}

/*
goos: windows
goarch: amd64
pkg: github.com/oldthreefeng/algo/leetcode
BenchmarkShortestPalindrome-4    	 5939828	       206 ns/op
BenchmarkShortestPalindromeX-4   	 3418819	       371 ns/op
*/