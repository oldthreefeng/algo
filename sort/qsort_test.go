/*
@Time : 2019/9/9 10:23
@Author : louis
@File : qsort_test
@Software: GoLand
*/

package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{32, 1, 7, -3, 10, 16, 8, -25, 2, 100}
	QuickSort(arr,0, len(arr)-1)
	//QuickSort1(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func TestQuickSort1(t *testing.T) {
	arr := []int{32, 1, 7, -3, 10, 16, 8, -25, 2, 100}
	//QuickSort(0, len(arr)-1, arr)
	QuickSort1(arr, 0, len(arr)-1)
	fmt.Println(arr)
}