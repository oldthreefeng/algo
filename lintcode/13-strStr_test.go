/*
@Time : 2019/9/26 16:15
@Author : louis
@File : 13-strStr_test
@Software: GoLand
*/

package lintcode

import (
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
