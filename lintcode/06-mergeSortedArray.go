/*
@Time : 2019/9/26 13:48
@Author : louis
@File : 06-mergeSortedArray
@Software: GoLand
*/

package lintcode

import "gogs.wangke.co/go/algo/sort"

/*
合并排序数组 II
中文English
合并两个有序升序的整数数组A和B变成一个新的数组。新数组也要有序。

Example
样例 1:

输入: A=[1], B=[1]
输出:[1,1]
样例解释: 返回合并后的数组。
样例 2:

输入: A=[1,2,3,4], B=[2,4,5,6]
输出: [1,2,2,3,4,4,5,6]
样例解释: 返回合并后的数组。
Challenge
你能否优化你的算法，如果其中一个数组很大而另一个数组很小？
*/

//a,b必须已经是排序的的
func MergeSortedArray(a, b []int) []int {
	var i, j, k int
	la, lb := len(a), len(b)
	lc := la + lb
	c := make([]int, lc)
	for i < la && j < lb {
		if a[i] < b[j] {
			c[k] = a[i]
			k++
			i++
		} else {
			c[k] = b[j]
			k++
			j++
		}
	}
	for i < la {
		c[k] = a[i]
		k++
		i++
	}
	for j < lb {
		c[k] = b[j]
		k++
		j++
	}
	return c
}

func MergeArrayX(a, b []int) []int{
	c := append(a,b...)
	sort.MergeSort(c,0,len(c)-1)
	return c
}
