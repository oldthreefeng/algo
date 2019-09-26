/*
@Time : 2019/9/26 17:23
@Author : louis
@File : 14-binarysearch_test
@Software: GoLand
*/

package lintcode

import "testing"

func TestBinarySearchX(t *testing.T) {
	a := []int{3,4,5,8,8,8,8,10,13,14}
	target := 8
	want := 3
	rel := BinarySearchX(a,target)
	if want != rel {
		t.Fatalf("want=%d, real=%d", want, rel)
	}
	t.Logf("want=%d", want)
}
