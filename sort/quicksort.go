package main

import (
	"fmt"
)

func QuickSort(left, right int, arr []int) {
	l, r := left, right
	pivot := arr[(left+right)/2]
	for l < r {
		//从pivot的左边找到大于pivot的值,小于向右移动指针
		for arr[l] < pivot {
			l++
		}
		//从pivot的右边找到小于pivot的值,大于向左移动指针
		for arr[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		//交换
		arr[l], arr[r] = arr[r], arr[l]
		// 优化,如果等于中轴值,相等
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	//向左递归,如果仅大于零,存在r和l相等的情况,就递归在里面了
	if r-left > 1 {
		QuickSort(left, r, arr)
	}
	//向右递归
	if right-l > 1 {
		QuickSort(l, right, arr)
	}
}

func main() {
	arr := []int{32, 1, 7, -3, 10, 16, 8, -25, 2, 100}
	//QuickSort(0, len(arr)-1, arr)
	QuickSort1(arr, 0, len(arr)-1)
	fmt.Println(arr)

}
func Swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func QuickSort1(a []int, s, e int) {
	if s >= e { //start end
		return
	}
	//起点位置
	k := a[s]
	var i, j = s, e
	for i != j {
		for j > i && a[j] >= k { //偶数次比较
			j--
		}
		Swap(a, i, j)
		for i < j && a[i] <= k { //奇数次比较
			i++
		}
		Swap(a, i, j)
	}
	QuickSort1(a, s, i-1)
	QuickSort1(a, i+1, e)

}
