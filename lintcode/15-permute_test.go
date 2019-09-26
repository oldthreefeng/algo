/*
@Time : 2019/9/26 18:01
@Author : louis
@File : 15-permute_test
@Software: GoLand
*/

package lintcode

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	a := []int{1,2,3}
	want := [][]int{
		{1,2,3},
		{3,2,1},
		{2,1,3},
		{1,3,2},
		{2,3,1},
		{3,1,2},
	}
	rel := Permute(&a)
	fmt.Println(want)
	fmt.Println(rel)
}
