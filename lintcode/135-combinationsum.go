/*
Copyright 2019 louis.
@Time : 2019/9/27 23:36
@Author : louis
@File : 135-combinationsum
@Software: GoLand

*/

package lintcode

import (
	"sort"
)

/*
### **Description**

中文English

给定一个候选数字的集合 `candidates` 和一个目标值 `target`. 找到 `candidates` 中所有的和为 `target` 的组合.

在同一个组合中, `candidates` 中的某个数字不限次数地出现.

所有数值 (包括 `target` ) 都是正整数.返回的每一个组合内的数字必须是非降序的.返回的所有组合之间可以是任意顺序.解集不能包含重复的组合.

Have you met this question in a real interview?  Yes

Problem Correction

### **Example**

**样例 1:**

```
输入: candidates = [2, 3, 6, 7], target = 7
输出: [[7], [2, 2, 3]]
```

**样例 2:**

```
输入: candidates = [1], target = 3
输出: [[1, 1, 1]]
```
*/



var (
	res [][]int = make([][]int, 0)
)

func CombinationSum(candidates []int, target int) [][]int {
	if candidates == nil || len(candidates) == 0 {
		return res
	}
	sort.Ints(candidates)
	l := make([]int, 0)
	dfsSum(candidates, l, 0, 0, target)
	return res
}

func dfsSum(candidates, l []int, s, val, target int) {
	if target == val {
		l1 := make([]int,len(l))
		copy(l1,l)   //必须进行值拷贝(深拷贝),引用拷贝(浅拷贝)无法正确运行
		//copy为值复制，更改原切片不会影响新切片，而等号复制相反
		res = append(res, l1)
		return
	}
	for i := s; i < len(candidates); i++ {
		if val + candidates[i] > target {
			return
		}
		l = append(l, candidates[i])
		dfsSum(candidates, l, i,val+candidates[i],target)
		l = l[:len(l)-1]
		for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
			i++
		}
	}

}
func CombinationSumX(candidates []int, target int) [][]int {
	if candidates == nil || len(candidates) == 0 {
		return res
	}
	sort.Ints(candidates)
	l := make([]int, 0)
	bfsSum(candidates, l, 0, 0, target)
	return res
}

func bfsSum(candidates, l []int, s, val, target int) {
	if target == val {
		l1 := make([]int,len(l))
		copy(l1,l)   //必须进行值拷贝(深拷贝),引用拷贝(浅拷贝)无法正确运行
		//copy为值复制，更改原切片不会影响新切片，而等号复制相反
		// slice的实现是底层数组的指针拷贝.
		res = append(res, l1)
		return
	}
	for i := s; i < len(candidates); i++ {
		if val + candidates[i] > target {
			return
		}
		l = append(l, candidates[i])
		bfsSum(candidates, l, i+1,val+candidates[i],target)
		l = l[:len(l)-1]
		for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
			i++
		}
	}

}