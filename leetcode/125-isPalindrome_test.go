/*
Copyright 2019 louis.
@Time : 2019/10/9 21:04
@Author : louis
@File : 125-isPalindrome
@Software: GoLand

*/

package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test01",args{"0P"},false},
		{"test02",args{"A man, a plan, a canal: Panama"},true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.s); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	s := "A man, a plan, a canal: Panama"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindrome(s)
	}
}
func BenchmarkIsPalindromeX(b *testing.B) {
	s := "A man, a plan, a canal: Panama"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindromeX(s)
	}
}

/*
BenchmarkIsPalindrome-4          6275407               191 ns/op             128 B/op          1 allocs/op
BenchmarkIsPalindromeX-4        12371260                95.8 ns/op             0 B/op          0 allocs/op

*/