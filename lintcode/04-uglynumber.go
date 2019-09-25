/*
@Time : 2019/9/25 18:24
@Author : louis
@File : 04-uglynumber
@Software: GoLand
*/

package lintcode

import "gogs.wangke.co/go/algo/utils"

/*
设计一个算法，找出只含素因子2，3，5 的第 n 小的数。

符合条件的数如：1, 2, 3, 4, 5, 6, 8, 9, 10, 12...

Example
样例 1：

输入：9
输出：10
样例 2：

输入：1
输出：1
Challenge
要求时间复杂度为 O(nlogn) 或者 O(n)。

Notice
我们可以认为 1 也是一个丑数。
*/
func UglyNumber(n int) int {
	if n <= 0 {
		return 0
	}
	var arr []int
	arr = append(arr, 1)
	var a, b, c int
	for len(arr) < n {
		//一个丑数都可以由之前的丑数乘（2,3,5)
		t1 := arr[a] * 2
		t2 := arr[b] * 3
		t3 := arr[c] * 5
		min := utils.Min(t1, utils.Min(t2, t3))
		arr = append(arr, min)
		if t1 == min {a++}
		if t2 == min {b++}
		if t3 == min {c++}
	}
	return arr[n-1]
}
