/*
@Time : 2019/9/26 17:22
@Author : louis
@File : 14-binarysearch
@Software: GoLand
*/

package lintcode

func BinarySearchX(nums []int, target int) int {
	// write your code here
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}
