/*
@Time : 2019/9/27 12:54
@Author : louis
@File : 29-partitionAarry_test
@Software: GoLand
*/

package lintcode

import (
	"testing"
)

func TestPartitionArray(t *testing.T) {
	var a []int
	rel := PartitionArray(a,10)
	t.Log(rel)

	testClass := []struct{
		a []int
		k int
		w int
	}{
		{[]int{9,9,9,8,9,8,7,9,8,8,8,9,8,9,8,8,6,9},9,10},
	}
	for _,tc := range testClass {
		rel := PartitionArray(tc.a,tc.k)
		want := tc.w
		if want != rel {
			t.Fatalf("want=%v, real=%v", want,rel)
		}
		t.Logf("want=%v", want)
	}
}
