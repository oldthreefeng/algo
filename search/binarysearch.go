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
