/*
@Time : 2019/9/17 23:12
@Author : louis
@File : maxsum
@Software: GoLand
*/

package main

import (
	"fmt"
	"gogs.wangke.co/go/algo/recursion"
)

/*
- 动态规划
问题:
{7},
{3, 8},
{8, 1, 0},
{2, 7, 4, 4},
{4, 5, 2, 6, 5},
在上面的数字三角形中寻找一条从顶部到底边的路径，使得路径上所经过的数字之和最大。
路径上的每一步都只能往左下或右下走。只需要求出这个最大和即可，不必给出具体路径。
三角形的行数大于1小于等于100，数字为0 -99

思路:
用二维数组存放数字三角形。
D( r, j)   : 第r行第j 个数字(r,j从1开始算)
MaxSum(r, j) :   从D(r,j)到底边的各条路径中，最佳路径的数字之和。
问题：求MaxSum(1,1)
典型的递归问题。
D(r, j)出发，下一步只能走D(r+1,j)或者D(r+1, j+1)。
故对于N行的三角形：
if  r == N    {
	MaxSum(r,j) = D(r,j)
} else {
	MaxSum( r, j) = Max{ MaxSum(r＋1,j), MaxSum(r+1,j+1) } + D(r,j)
}
*/

const MAX = 5

var (
	//存储直角三角形的二维数组
	D      [MAX][MAX]int
	//二维数组的列数
	n      int
	//递推时存入将和存入一维
	MaxSum [MAX]int
	//递归防止重复计算
	maxSum [MAX][MAX]int
)


//此方法会超时,O(2^n)时间复杂度,需要将算出的每一个值进行保留
func MaxSumRe(i, j int) int {
	if i == n {
		return D[i][j]
	}
	x := MaxSumRe(i+1, j)
	y := MaxSumRe(i+1, j+1)
	return recursion.Max(x, y) + D[i][j]
}

//基于MaxSumRe优化,将计算的结果存入二维数组.
func MaxSumReC(i, j int) int {
	if maxSum[i][j] != -1 {
		return maxSum[i][j]
	}
	if i == n {
		maxSum[i][j] = D[i][j]
	} else {
		x := MaxSumRe(i+1, j)
		y := MaxSumRe(i+1, j+1)
		maxSum[i][j] = D[i][j] + recursion.Max(x, y)
	}
	return maxSum[i][j]
}

func MaxSumFor() {
	MaxSum = D[n]
	//将最大路径数存入到D二维数组的最后一个
	for i := n-1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			MaxSum[j] = recursion.Max(MaxSum[j], MaxSum[j+1]) + D[i][j]
		}
	}
	fmt.Println(MaxSum[0])
}



func main() {
	//使用递归解决
	fmt.Scanln(&n)
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Scanln(&D[i][j])
			maxSum[i][j] = -1
		}
	}
	fmt.Println(MaxSumReC(0, 0))

	//2. 使用递推解决
	fmt.Scanln(&n)
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Scanln(&D[i][j])
			//maxSum[i][j] = -1
		}
	}
	//fmt.Println(MaxSumReC(0, 0))
	MaxSumFor()
}

/*
	D = [MAX][MAX]int{
		{7},
		{3, 8},
		{8, 1, 0},
		{2, 7, 4, 4},
		{4, 5, 2, 6, 5},
	}
*/
