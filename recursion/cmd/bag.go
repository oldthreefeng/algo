/*
@Time : 2019/9/18 20:43
@Author : louis
@File : bag
@Software: GoLand
*/

package main

import (
	"fmt"
	"github.com/oldthreefeng/algo/utils"
)

/*
问题描述:

	有N件物品和一个容积为M的背包。第i件物品的体积w[i]，价值是d[i]。
	求解将哪些物品装入背包可使价值总和最大。
	每种物品只有一件，可以选择放或者不放(N<=3500,M <= 13000)。

思路:

	用F[i][ j] 表示取前i种物品，使它们总体积不超过j的最优取法取得的价值总和。要求F[N][M]
	边界条件:
	if w[1] <= j {
		F[1][j] = d[1]
	} else {
		F[1][j] = 0
	}
	if j-w[i] >= 0 {
		//分两步走,取i-1种物品的最优解 + 不取i-1物品的最优解
		F[i][j] = max { F[i-i][j], F[i-1][j-w[i]] +d[i]}
	} else {
		F[i][j] = F[i-i][j]
	}


思考:

	本题如用记忆型递归，需要一个很大的二维数组，会超内存。
	注意到这个二维数组的下一行的值，只用到了上一行的正上方及左边的值，
	因此可用滚动数组的思想，只要一行即可。即可以用一维数组，用“人人为我”递推型动归实现。

*/

func main() {
	Bag()
	BagG()
}

func Bag()  {
	const (
		n   = 5
		bag = 15
	)
	w := [n]int{5, 4, 7, 2, 6}
	d := [n]int{12, 3, 10, 3, 6}
	var f [n][bag + 1]int
	for i := 1; i < n; i++ {
		for j := 1; j <= bag; j++ {
			if j-w[i] >= 0 {
				f[i][j] = utils.Max(f[i-1][j], f[i-1][j-w[i]]+d[i])
			} else {
				f[i][j] = f[i-1][j]
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= bag; j++ {
			fmt.Print(f[i][j], " ")
		}
		fmt.Println()
	}

}

func BagG() {
	const (
		n   = 5
		bag = 15
	)
	w := [n]int{5, 4, 7, 2, 6}
	d := [n]int{12, 3, 10, 3, 6}
	var f [bag+1]int

	for i := 1; i < n; i++ {
		for j := bag; j >= w[i]; j-- {
			//时间无法优化,空间进行了优化.
			f[j] = utils.Max(f[j], f[j-w[i]]+d[i])
		}
	}
	fmt.Println(f)
}
