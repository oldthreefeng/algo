package main

import "fmt"

func swap(s *[5]int, i, j int) {
	s[i], s[j] = s[j], s[i]
}

func insectionSort(arr *[5]int) {
	lens := len(arr)
	for i := 1; i < lens; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			// 取出下一个元素，在已经排序的元素序列中从后向前扫描，
			// 如果该元素大于新元素，将该元素移到下一位置；把arr[j]和arr[j-1]交换,
			swap(arr, j, j-1)
		}
	}
}

func main() {
	arr := [5]int{11, 3, 98, 76, 32}
	fmt.Println(arr)
	insectionSort(&arr)
	fmt.Println(arr)

}
