/*
@Time : 2019/9/26 17:37
@Author : louis
@File : 15-permute
@Software: GoLand
*/

package lintcode

//This is NOT WORK
//TODO

func Permute(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) <= 0 {
		return res
	}
	cur := make([]int, len(nums))
	used := make([]bool, len(nums))
	dfs(res, cur, nums, used)
	return res
}

func dfs(res [][]int, cur []int, nums []int, used []bool) {
	if len(cur) == len(nums) {
		var cur1 = make([]int, 0)
		copy(cur1, cur)
		res = append(res, cur)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !used[i]  {
			cur = append(cur, nums[i])
			used[i] = true
			dfs(res, cur, nums, used)
			used[i] = false
			cur = cur[:len(cur)-1]
		}
	}
}

func pop(a []int, n int) []int {
	j := 0
	for _, val := range a {
		if val == n {
			a[j] = val
			j++
		}
	}
	return a[:j]
}
