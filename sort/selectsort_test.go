/*
@Time : 2019/9/9 10:26
@Author : louis
@File : selectsort_test
@Software: GoLand
*/

package sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	arr := [5]int{11, 3, 98, 76, 32}
	fmt.Println(arr)
	SelectSort(&arr)
	fmt.Println(arr)
}
