/*
@Time : 2019/9/27 11:14
@Author : louis
@File : 28-searchMatrix_test
@Software: GoLand
*/

package lintcode

/*
### Description

写出一个高效的算法来搜索 *m* × *n*矩阵中的值。

这个矩阵具有以下特性：

- 每行中的整数从左到右是排序的。
- 每行的第一个数大于上一行的最后一个整数。

### **Example**

```
样例  1:
	输入: [[5]],2
	输出: false

	样例解释:
  没有包含，返回false。

样例 2:
	输入:
[
  [1, 3, 5, 7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
],3
	输出: true

	样例解释:
	包含则返回true。
```

### **Challenge**

O(log(n) + log(m)) 时间复杂度
*/

import (
	"testing"
)

func TestSearchMatrix(t *testing.T) {

	testClass := []struct{
		m [][]int
		tg int
		b bool
	}{
		{[][]int{{1,3,5,7}, {10,11,16,20},{23,30,34,50}},3,true},
		{[][]int{{5}},3,false},
		{[][]int{{5,7,10}},10,true},
		{[][]int{{}},10,false},
		{nil,10,false},
	}
	for _,tc := range testClass {
		want := tc.b
		rel := SearchMatrix(tc.m,tc.tg)
		if want != rel {
			t.Fatalf("want=%v, real=%v", want,rel)
		}
		t.Logf("want=%v", want)
	}
}
