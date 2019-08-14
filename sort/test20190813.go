package main

import (
	"fmt"
	"sort"
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
func Test04(left, right int, arr *[num]int) {
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

	//向左递归
	if r-left > 1 {
		Test04(left, r, arr)
	}
	//向右递归
	if right-l > 1 {
		Test04(l, right, arr)
	}
}

func Test05(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	Test05(data[:head])
	Test05(data[head+1:])
}

func main() {
	arr := [num]int{-9, 78, 0, 23, -567, 70, 123, 90, -23, 10}
	fmt.Println("初始", arr)
	//test01(arr)
	//test03(arr)
	//test02(arr)
	//Test04(0, num-1, arr)
	//切片直接指针操作
	arrS := []int{-9, 78, 0, 23, -567, 70, 123, 90, -23, 10}
	//Test05(arrS)
	fmt.Println("main..")
	sort.Ints(arrS)
	fmt.Println(arrS)
	//fmt.Println(arr)

}
