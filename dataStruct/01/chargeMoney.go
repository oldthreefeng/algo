/*
@Time : 2019/9/2 22:06
@Author : louis
@File : chargeMoney
@Software: GoLand
*/

package main

import "fmt"

//笨办法,直接三重循环O(n^3)
func cha1() {
	for i := 0; i <= 50; i++ {
		for j := 0; j <= 50; j++ {
			for k := 0; k <= 50; k++ {
				if i+k+j == 50 && i+2*j+5*k == 100 {
					fmt.Printf("%d,%d,%d\n", i, j, k)
				}
			}
		}
	}
}

//改进办法,直接二重循环O(n^2)
func cha2() {
	for i := 0; i <= 50; i++ {
		for j := 0; j <= 50; j++ {
			k := 50 - i - j
			if i+2*j+5*k == 100 {
				fmt.Printf("%d,%d,%d\n", i, j, k)
			}

		}
	}
}

// 算法复杂度O(1),利用数学公式
func cha3() {
	for k := 0; k < 13; k++ {
		j := 50 - 4*k
		i := 50 - j - k
		fmt.Printf("%d,%d,%d\n", i, j, k)
	}
}

//设计求解下列问题的算法，并分析其最坏情况的时间复杂度及其量级。
//
//（1）在数组A[1..n]中查找值为K的元素，若找到则输出其位置i(1<=i<=n)，否则输出0作为标志。
func findK(arr []int, K int) int {
	for k, v := range arr {
		if v == K {
			return k + 1
		}
	}
	return 0
}

//
//（2）找出数组A[1..n]中元素的最大值和次最大值（本小题以数组元素的比较为标准操作）。
func findM(arr []int) (maxIndex, secIndex int) {
	var max, sec int
	if arr[0] < arr[1] {
		max, sec = arr[1], arr[0]
		maxIndex, secIndex = 1, 0
	} else {
		max, sec = arr[0], arr[1]
		maxIndex, secIndex = 0, 1
	}
	for k, v := range arr {
		if max < v {
			max = v
			maxIndex = k
		}
	}
	for k, v := range arr {
		if sec < v && k != maxIndex {
			sec = v
			secIndex = k
		}
	}

	return maxIndex, secIndex
}

func main() {
	//cha1()
	//cha2()
	//cha3()

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(findK(arr, 8))
	//arr = []int{-1, -2, -3, -4, -6, -80, -5, -6, -7, -8, -9, -10}
	fmt.Println(findM(arr))
}
