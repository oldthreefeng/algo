/*
@Time : 2019/9/27 12:54
@Author : louis
@File : 29-partitionArray
@Software: GoLand
*/

package lintcode

/*
快排思想
将小于k的全部放到k的左边
将大于k的全部放到k的右边
*/

func PartitionArray(nums []int, k int) int {
	// write your code here
	if nums == nil || len(nums) == 0 || nums[0] == 0 {
		return 0
	}
	s, e := 0, len(nums)-1
	for s < e {
		for s <= e && nums[s] < k {
			s++
		}
		for s <= e && nums[e] >= k {
			e--
		}
		if s <= e {
			nums[s], nums[e] = nums[e], nums[s]
			s++
			e--
		}
	}

	return s
}
