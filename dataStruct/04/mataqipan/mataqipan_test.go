/*
@Time : 2019/9/27 9:59
@Author : louis
@File : mataqipan_test
@Software: GoLand
*/

package mataqipan

import (
	"fmt"
	"testing"
)

// 全局变量导致不可用
func TestSolution(t *testing.T) {
	want :=0
	testClass := []struct{
		x int
		y int
	}{
		{0,0},{0,1},{0,2},{0,3},
		{1,0},{1,1},{1,2},{1,3},
		{2,0},{2,1},{2,2},{2,3},
		{3,0},{3,1},{3,2},{3,3},
		{4,0},{4,1},{4,2},{4,3},
	}
	for _,tc :=range testClass {
		rel := Solution(tc.x,tc.y)
		if want!=rel {
			fmt.Println(rel)
		}
	}
}
