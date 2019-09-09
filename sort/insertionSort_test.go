/*
@Time : 2019/9/9 10:20
@Author : louis
@File : insertionSort_test
@Software: GoLand
*/

package sort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{11, 3, 98, 76, 32}
	fmt.Println(arr)
	InsertionSort(arr)
	//UpInsertionSort(&arr)
	fmt.Println(arr)

}

func TestUpInsertionSort(t *testing.T) {
	arr := []int{11, 3, 98, 76, 32}
	fmt.Println(arr)
	//InsertionSort(&arr)
	UpInsertionSort(arr)
	fmt.Println(arr)
}
