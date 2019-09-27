/*
Copyright 2019 louis.
@Time : 2019/9/26 22:33
@Author : louis
@File : 08-rotateString_test
@Software: GoLand

*/

package lintcode

import "testing"
/*
给定一个字符串（以字符数组的形式给出）和一个偏移量，根据偏移量原地旋转字符串(从左向右旋转)。

Example
样例 1:

输入:  str="abcdefg", offset = 3
输出:  str = "efgabcd"
样例解释:  注意是原地旋转，即str旋转后为"efgabcd"
样例 2:

输入: str="abcdefg", offset = 0
输出: str = "abcdefg"
样例解释: 注意是原地旋转，即str旋转后为"abcdefg"
样例 3:

输入: str="abcdefg", offset = 1
输出: str = "gabcdef"
样例解释: 注意是原地旋转，即str旋转后为"gabcdef"
样例 4:

输入: str="abcdefg", offset =2
输出: str = "fgabcde"
样例解释: 注意是原地旋转，即str旋转后为"fgabcde"
样例 5:

输入: str="abcdefg", offset = 10
输出: str = "efgabcd"
样例解释: 注意是原地旋转，即str旋转后为"efgabcd"
Challenge
在数组上原地旋转，使用O(1)的额外空间
*/

func TestRotateString(t *testing.T) {
	testClass := []struct{
		str string
		offset int
		rel string
	}{
		{"abcdefg",3,"efgabcd"},
		{"abcdefg",0,"abcdefg"},
		{"abcdefg",10,"efgabcd"},
		{"abcdefg",20,"bcdefga"},
	}
	for _,tc :=range  testClass {
		RotateString(&tc.str,tc.offset)
		want := tc.str
		rel := tc.rel
		if want != rel {
			t.Fatalf("want=%s, real=%s", want,rel)
		}
		t.Logf("want=%s", want)
	}
}
