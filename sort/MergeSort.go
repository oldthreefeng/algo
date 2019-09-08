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

import (
	"fmt"
	"math/rand"
	"time"
)

// 合并 [s,e] 两部分数据，m 左半部分的终点，m + 1 是右半部分的起点
func merge(arr []int, s int, m int, e int) {
	// 因为需要直接修改 arr 数据，这里首先复制 [s,e] 的数据到新的数组中，用于赋值操作
	// 归并的操作为O (e-m+1)
	temp := make([]int, e-s+1)
	for i := s; i <= e; i++ {
		temp[i-s] = arr[i]
	}

	// 指向两部分起点
	left := s
	right := m + 1

	for i := s; i <= e; i++ {
		// 左边的点超过终点点，说明只剩右边的数据
		if left > m {
			arr[i] = temp[right-s]
			right++
			// 右边的数据超过终点，说明只剩左边的数据
		} else if right > e {
			arr[i] = temp[left-s]
			left++
			// 左边的数据大于右边的数据，选小的数字
		} else if temp[left-s] > temp[right-s] {
			arr[i] = temp[right-s]
			right++
		} else {
			arr[i] = temp[left-s]
			left++
		}
	}
}

func MergeSort(arr []int, s int, e int) {
	if s >= e {
		return
	}

	// 递归向下
	mid := s + (e-s)/2
	MergeSort(arr, s, mid)
	MergeSort(arr, mid+1, e)
	// 归并向上
	merge(arr, s, mid, e)
}

func main() {
	//arr := []int{3, 1, 2, 5, 6, 43, 4}
	//ArrangeRight(arr,0, len(arr)-1,4)
	//MergeSort(arr, len(arr)-5, len(arr)-1)
	//fmt.Println(arr[len(arr)-4:])
	n := 100000
	m := 1000
	var arr []int
	for i:=0;i< n;i++ {
		//arr= append (arr, i) //刚好生成的有序数组,时间会超时
		arr = append(arr, rand.Intn(n))
	}
	start := time.Now().Unix()
	ArrangeRight(arr,0, n-1,m)
	MergeSort(arr, n-m, n-1)
	end := time.Now().Unix()
	fmt.Print(end-start)
}

func ArrangeRight(arr []int, s, e, k int) {
	if s >= e {
		return
	}
	if k == e-s+1 {
		return
	}
	key := arr[s]
	i, j := s, e
	//前a[s]大的全部排到右边
	for i != j {
		for j > i && arr[j] >= key {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
		for j > i && arr[i] <= key {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	// 如果当前的t[s]的位置等于k.前k大的已经排序完毕
	a := e - i + 1
	if k == a {
		return
	}else if k < a {
		ArrangeRight(arr, i+1, e, k)
	}else {
		ArrangeRight(arr, s, i-1, k-a)
	}

}
