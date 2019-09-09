/*
@Time : 2019/9/9 10:21
@Author : louis
@File : mergesort_test
@Software: GoLand
*/

package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{3, 1, 2, 5, 6, 43, 4}
	MergeSort(arr, len(arr)-5, len(arr)-1)
	fmt.Println(arr[len(arr)-4:])
}
