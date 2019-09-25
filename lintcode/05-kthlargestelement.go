/*
@Time : 2019/9/25 19:30
@Author : louis
@File : 05-kthlargestelement
@Software: GoLand
*/

package lintcode

import "gogs.wangke.co/go/algo/sort"

// O(n+mlogm) 超时
func KthLargestElement(n int, nums []int) int {
	sort.ArrangeRight(nums, 0, len(nums)-1, n)
	return nums[len(nums)-n]
}

func KetLargestElementX(n int, nums []int) int {
	return 0
}
