/*
@Time : 2019/9/18 20:25
@Author : louis
@File : ways
@Software: GoLand
*/

package main

import "fmt"

var (
	a [30]int
	ways [40][30]int
	N int
)

func Ways(i, j int) int {
	//从前j种物品选择,凑成体积为i的数目
	if i == 0 {
		return 1
	}
	if j <= 0 { //如果物品j小于等于0,说明没有
		return 0
	}
	return Ways(i, j-1) + Ways(i-a[j], j-1)
}

func main() {
	//递归方法
	fmt.Scanln(&N)
	for i := 1; i <= N; i++ {
		fmt.Scanln(&a[i])
		ways[0][i] = 1
	}
	fmt.Println(Ways(40, N))

	// 动态规划方法
	ways[0][0]=1
	for i := 1; i <= 40; i++ {
		for j := 1; j <= N; j++ {
			ways[i][j] = ways[i][j-1]
			if i-a[j]>=0 {
				ways[i][j] += ways[i-a[j]][j-1]
			}
		}
	}
	fmt.Println(ways[39][N])
}
