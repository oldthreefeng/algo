/**
 * @time: 2019-08-20 22:24
 * @author: louis
 */
package recursion

import (
	"fmt"
	"math"
)


func Count24(nums []int) bool {
	arr := make([]float64, len(nums))
	for i := range nums {
		arr[i] = float64(nums[i])
	}
	return dfs(arr)
}

func dfs(nums []float64) bool {
	if len(nums) == 1 {
		return math.Abs(nums[0] - 24) < 0.001
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			next := []float64{}
			for k := 0; k < len(nums); k++ {
				if k != i && k != j {
					next = append(next, nums[k])
				}
			}

			l := len(next)
			for _, k := range []byte{'+', '-', '*', '/'} {
				var tmp float64
				switch k {
				case '+':
					tmp = nums[i] + nums[j]
				case '-':
					tmp = nums[i] - nums[j]
				case '*':
					tmp = nums[i] * nums[j]
				case '/':
					tmp = nums[i] / nums[j]
				}
				next = append(next, tmp)
				ok := dfs(next)
				if ok {
					return true
				}
				next = next[:l]
			}
		}
	}
	return false
}



func JudgePoint24(nums []int) bool {
	arr := make([]float64, 4)
	for i := 0; i < 4; i++ {
		arr[i] = float64(nums[i])
	}
	return backtrace(arr)
}

func backtrace(nums []float64) bool {
	zero := 0.0001
	if len(nums) == 1 {
		if math.Abs(nums[0]-float64(24)) < zero {
			return true
		}
		return false
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			p := nums[i]
			q := nums[j]
			add := []float64{p + q, p - q, q - p, p * q}
			if math.Abs(p) > zero {
				add = append(add, q/p)
			}
			if math.Abs(q) > zero {
				add = append(add, p/q)
			}

			nums = append(nums[:i], nums[i+1:]...) // 移除i元素
			nums = append(nums[:j], nums[j+1:]...) // 移除j元素

			for _, v := range add {
				nums = append(nums, v)
				fmt.Println(nums)
				if backtrace(nums) {
					return true
				}
				nums = nums[:len(nums)-1]
			}

			rear := append([]float64{}, nums[j:]...)
			nums = append(nums[:j], q)
			nums = append(nums, rear...)

			rear = append([]float64{}, nums[i:]...)
			nums = append(nums[:i], p)
			nums = append(nums, rear...)
		}
	}
	return false
}