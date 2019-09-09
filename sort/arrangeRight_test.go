/*
@Time : 2019/9/9 10:15
@Author : louis
@File : arrangeRight_test
@Software: GoLand
*/

package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestArrangeRight(t *testing.T) {
	n := 1000000
	m := 1000
	var arr []int
	for i:=0;i< n;i++ {
		//arr= append (arr, i) //刚好生成的有序数组,时间会超时
		arr = append(arr, rand.Intn(n))
	}
	start := time.Now().Unix()
	ArrangeRight(arr,0, n-1,m)
	MergeSort(arr, n-m, n-1)
	end := time.Now().Unix()
	fmt.Print(end-start)
}
