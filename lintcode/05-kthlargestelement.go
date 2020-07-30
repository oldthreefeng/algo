/*
@Time : 2019/9/25 19:30
@Author : louis
@File : 05-kthlargestelement
@Software: GoLand
*/

package lintcode

import "github.com/oldthreefeng/algo/sort"

// O(n+mlogm) 超时,但是不稳定,最差的是O(n^2)
func KthLargestElement(n int, nums []int) int {
	sort.ArrangeRight(nums, 0, len(nums)-1, n)
	return nums[len(nums)-n]
}

//O(nlogn) 稳定
func KetLargestElementX(n int, nums []int) int {
	sort.MergeSort(nums,0,len(nums)-1)
	return nums[len(nums)-n]
}
