package main

import (
	"fmt"
)

const num = 10

func test01(arr *[num]int) {
	for i := 0; i < num-1; i++ {
		for j := 1; j < num; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

func test02(arr *[num]int) {
	for i := 0; i < num-1; i++ {
		min := i
		for j := i + 1; j < num; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func test03(arr *[num]int) {
	for i := 1; i < num; i++ {
		insertVal := arr[i] //从第二个元素开始比较.
		var insert int
		for insert = i - 1; insert >= 0 && arr[insert] > insertVal; insert-- {
			arr[insert+1] = arr[insert]
		}
		arr[insert+1] = insertVal //如果insert == i-1,没有进行比较,赋值依旧可以
	}
}

//
//func QuickSort(left, right int, arr *[num]int) {
//	l, r := left, right
//	pivot := arr[(left+right)/2]
//	for l < r {
//		//从pivot的左边找到大于pivot的值
//		for arr[l] < pivot {
//			l++
//		}
//		//从pivot的右边找到小于pivot的值
//		for arr[r] > pivot {
//			r--
//		}
//		if l >= r {
//			break
//		}
//		//交换
//		arr[l], arr[r] = arr[r], arr[l]
//		// 优化
//		if arr[l] == pivot {
//			r--
//		}
//		if arr[r] == pivot {
//			l++
//		}
//	}
//	//如果l和r相等,就错位
//	if l == r {
//		l--
//		r++
//	}
//	//向左递归
//	if left < r {
//		QuickSort(left, r, arr)
//	}
//	//向右递归
//	if right > l {
//		QuickSort(l, right, arr)
//	}
//}

//func main() {
//	arr := &[num]int{1, -20, 31, 400, 0, 16, 7, 28, -39, 10}
//	//test01(arr)
//	//test03(arr)
//	//test02(arr)
//	QuickSort(0, 9, arr)
//	fmt.Println(arr)
//}

func QuickSort(left int, right int, array *[10]int) {
	l := left
	r := right
	// pivot 是中轴， 支点
	pivot := array[(left+right)/2]
	temp := 0
	//for 循环的目标是将比 pivot 小的数放到 左边
	// 比 pivot 大的数放到 右边
	for l < r {
		//从 pivot 的左边找到大于等于 pivot 的值
		for array[l] < pivot {
			l++
		}
		//从 pivot 的右边边找到小于等于 pivot 的值
		for array[r] > pivot {
			r--
		}
		// 1 >= r 表明本次分解任务完成, break
		if l >= r {
			break
		}
		//交换
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		//优化
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	// 如果 1== r, 再移动下
	if l == r {
		l++
		r--
	}
	// 向左递归
	if left < r {
		QuickSort(left, r, array)
	}
	// 向右递归
	if right > l {
		QuickSort(l, right, array)
	}
}

func main() {
	arr := [10]int{-9, 78, 0, 23, -567, 70, 123, 90, -23, 10}
	fmt.Println("初始", arr)
	//调用快速排序
	QuickSort(0, len(arr)-1, &arr)
	fmt.Println("main..")
	fmt.Println(arr)
}
