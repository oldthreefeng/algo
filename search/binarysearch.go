package search

//BinarySearch 从小到大的int数组a里面,
func BinarySearch(arr []int, size, p int) int {
	L := 0
	R := size - 1
	for L < R {
		mid := L + (R-L)/2
		if p == mid {
			return mid
		} else if p > mid {
			L = mid + 1
		} else {
			R = mid - 1
		}
	}
	return -1
} // O(log n)

func BinarySearchX(nums []int, target int) int {
	// write your code here
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
