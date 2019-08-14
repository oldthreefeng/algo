package main

import (
	"bufio"
	"fmt"
	"os"
)

const num = 10

func QuickSort(left, right int, arr *[num]int) {
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
		QuickSort(left, r, arr)
	}
	//向右递归
	if right-l > 1 {
		QuickSort(l, right, arr)
	}
}

func main() {
	//
	////reader := bufio.NewReader(os.Stdin)
	////no,_ := reader.ReadBytes('\n')
	////n32 := binary.BigEndian.Uint32(no)
	////n:= int(n32)
	//var n int
	//fmt.Scanln(&n)
	//
	//var arr []int
	//for i := 0; i < n; i++ {
	//	//no,err := reader.ReadBytes('\n')
	//	//if err == io.EOF {
	//	//	break
	//	//}
	//	//j32 := binary.BigEndian.Uint32(no)
	//	//j := int(j32)
	//	var j int
	//	fmt.Scanln(&j)
	//	arr = append(arr, j)
	//}
	//
	//for i := 0; i < len(arr); i++ {
	//	for j := 0; j < len(arr)-i-1; j++ {
	//		if arr[j] > arr[j+1] {
	//			arr[j], arr[j+1] = arr[j+1], arr[j]
	//		}
	//	}
	//}
	//
	//i := 0
	//var j int
	//for {
	//	if i >= len(arr)-1 {
	//		break
	//	}
	//	for j = i + 1; j < len(arr) && arr[i] == arr[j]; j++ {
	//	}
	//	arr = append(arr[:i+1], arr[j:]...)
	//	i++
	//}
	//for _, v := range arr {
	//	fmt.Println(v)
	//}
	//for k := range sli {
	//	fmt.Println(k)
	//}

	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	for _,v := range str {
		defer fmt.Printf(string(v))
	}

}
