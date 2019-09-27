/*
@Time : 2019/9/26 16:15
@Author : louis
@File : 13-strStr_test
@Software: GoLand
*/

package lintcode

import (
	"strings"
	"testing"
)

func TestStrStr(t *testing.T) {
	source := "abcdabcdefg"
	target := "bcd"
	want := 1
	rel := StrStr(source, target)
	if want != rel {
		t.Fatalf("want=%d, real=%d", want, rel)
	}
	t.Logf("want=%d", want)
}

func BenchmarkStrStr(b *testing.B) {
	source := "abcdabcdefg"
	target := "bcd"
	b.ResetTimer()
	for i:=0;i<b.N;i++ {
		StrStr(source,target)
	}
}

func BenchmarkStrStr2(b *testing.B) {
	source := "abcdabcdefg"
	target := "bcd"
	b.ResetTimer()
	for i:=0;i<b.N;i++ {
		strings.Index(source,target)
	}
}

/*
BenchmarkStrStr-4    	84754094	        14.6 ns/op
BenchmarkStrStr2-4   	139050566	         8.72 ns/op
*/