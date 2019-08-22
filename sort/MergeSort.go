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

// 合并 [s,e] 两部分数据，m 左半部分的终点，m + 1 是右半部分的起点
func merge(arr []int, s int, m int, e int) {
	// 因为需要直接修改 arr 数据，这里首先复制 [s,e] 的数据到新的数组中，用于赋值操作
	temp := make([]int, e-s+1)
	for i := s; i <= e; i++ {
		temp[i-s] = arr[i]
	}

	// 指向两部分起点
	left := s
	right := m + 1

	for i := s; i <= e; i++ {
		// 左边的点超过中点，说明只剩右边的数据
		if left > m {
			arr[i] = temp[right-s]
			right++
			// 右边的数据超过终点，说明只剩左边的数据
		} else if right > e {
			arr[i] = temp[left-s]
			left++
			// 左边的数据大于右边的数据，选小的数字
		} else if temp[left - s] > temp[right - s] {
			arr[i] = temp[right - s]
			right++
		} else {
			arr[i] = temp[left - s]
			left++
		}
	}
}

func MergeSort(arr []int, s int, e int) {
	if s >= e {
		return
	}

	// 递归向下
	mid := s + (e-s) / 2
	MergeSort(arr, s, mid)
	MergeSort(arr, mid+1, e)
	// 归并向上
	merge(arr, s, mid, e)
}

func main() {
	arr := []int{3, 1, 2, 5, 6, 43, 4}
	MergeSort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}
