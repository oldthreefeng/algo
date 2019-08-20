/**
 * @time: 2019-08-20 23:24
 * @author: louis
 */

/*
   T(n) = 2*T(n/2) + a*n  (a是常数,具体多少不重要）
		= 2*(2*T(n/4)+a*n/2)+a*n
		= 4*T(n/4)+2a*n
		= 4*(2*T(n/8)+a*n/4)+2*a*n
		= 8*T(n/8)+3*a*n...
		= 2k*T(n/2k)+k*a*n
		一直做到n/2k= 1 (此时k = log2n)，
		T(n)= 2k*T(1)+k*a*n
		= 2k*T(1)+k*a*n = 2k+k*a*n
		= n+a*(log2n)*n   //O(n*log n)
*/
//func MergeSort(a []int, s, e int) {
//	if s < e {
//		var m = s + (e-s)/2
//		MergeSort(a, s, m)
//		MergeSort(a, m+1, e)
//
//		if a[m] > a[m+1] {
//			Merge(a, s, m, e)
//		}
//	}
//}
//
//func Merge(a []int, s, m, e int) {
//	tmp := make([]int, e-s+1)
//	for i := s; i <= e; i++ {
//		tmp[i-s] = a[i]
//	}
//
//	left := s
//	right := m + 1
//	for i := s; i <= e; i++ {
//		if left > m {
//			a[i] = tmp[right-s]
//			right++
//		} else if right > e {
//			a[i] = tmp[left-s]
//			left++
//		} else if tmp[left-s] > tmp[right-s] {
//			a[i] = tmp[right-s]
//		} else {
//			a[i] = tmp[left-s]
//			left++
//		}
//	}
//}
//
//func main() {
//	arr := []int{3, 1, 2, 5, 6, 43, 4}
//	MergeSort(arr, 0, len(arr)-1)
//
//	fmt.Println(arr)
//}

package main

import "fmt"

// 合并 [l,r] 两部分数据，mid 左半部分的终点，mid + 1 是右半部分的起点
func merge(arr []int, l int, mid int, r int) {
	// 因为需要直接修改 arr 数据，这里首先复制 [l,r] 的数据到新的数组中，用于赋值操作
	temp := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		temp[i-l] = arr[i]
	}

	// 指向两部分起点
	left := l
	right := mid + 1

	for i := l; i <= r; i++ {
		// 左边的点超过中点，说明只剩右边的数据
		if left > mid {
			arr[i] = temp[right-l]
			right++
			// 右边的数据超过终点，说明只剩左边的数据
		} else if right > r {
			arr[i] = temp[left-l]
			left++
			// 左边的数据大于右边的数据，选小的数字
		} else if temp[left - l] > temp[right - l] {
			arr[i] = temp[right - l]
			right++
		} else {
			arr[i] = temp[left - l]
			left++
		}
	}
}

func MergeSort(arr []int, l int, r int) {
	if l >= r {
		return
	}

	// 递归向下
	mid := (r + l) / 2
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)
	// 归并向上
	merge(arr, l, mid, r)
}

func main() {
	arr := []int{3, 1, 2, 5, 6, 43, 4}
	MergeSort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}
