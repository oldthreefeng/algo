package main

import "fmt"

func BubbleSort(arr *[5]int) {
	fmt.Println("before sort:", arr)
	lens := len(*arr)
	tmp := 0
	for j := 0; j < lens-1; j++ {
		for i := 0; i < lens-1-j; i++ {
			if arr[i] > arr[i+1] {
				tmp = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
			}
		}
	}
	fmt.Println("sorted :", arr)

}

func main() {
	//定义数组
	arr := [5]int{24, 69, 80, 57, 13}
	BubbleSort(&arr)
	fmt.Println("main arr:", arr)
}
