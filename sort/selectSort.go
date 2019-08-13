package main

import "fmt"

func selectSort(arr *[5]int) *[5]int {
	// *取值符 ， &取址符
	lens := len(arr)
	for i := 0; i < lens-1; i++ {
		minIndex := i
		for j := i + 1; j < lens; j++ {
			if arr[j] < arr[minIndex] { //找到最小的元素下标
				minIndex = j
			}
		}
		tmp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = tmp
	}
	return arr
}

func main() {
	arr := [5]int{11, 3, 98, 76, 32}
	fmt.Println(arr)
	selectSort(&arr)
	fmt.Println(arr)

}

//TODO  other Algorithm in sort
