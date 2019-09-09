/*
@Time : 2019/9/9 10:18
@Author : louis
@File : bubblesort_test
@Software: GoLand
*/

package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{24, 69, 80, 57, 13}
	BubbleSort(arr)
	fmt.Println("main arr:", arr)
}
