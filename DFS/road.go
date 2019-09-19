/*
@Time : 2019/9/19 0:25
@Author : louis
@File : road
@Software: GoLand
*/

package main

import (
	"fmt"
	"gogs.wangke.co/go/algo/utils"
)

/*
[百炼 1724](http://bailian.openjudge.cn/practice/1724/)

N个城市，编号1到N。
城市间有R条单向道路。每条道路连接两个城市，有长度和过路费两个属性。
Bob只有K块钱，他想从城市1走到城市N。
问最短共需要走多长的路。如果到不了N，输出-1
2<=N<=100
0<=K<=10000
1<=R<=10000
每条路的长度L, 1 <= L <= 100
每条路的过路费T , 0 <= T <= 100
输入：
K
N
R
s1 e1 L1 T1
s1 e2 L2 T2...
sR eR LR TR
s e是路起点和终点

样例输入:
5
6
7
1 2 2 3
2 4 3 3
3 4 2 4
1 3 4 1
4 6 2 1
3 5 2 0
5 4 3 2

输出:
11

*/

var (
	N         int             //N个城市，编号1到N
	R         int             //城市间有R条单向道路
	K         int             //有K块钱
	cityMap   [110][110]Road  //采用二维数组存储图,cityMap[i]是从点i有路连接到城市的集合
	minLen    int             //到目前为止,从1到n的路最佳走法的长度.
	totalLen  int             //目前在走的路径长度
	totalCost int             //目前再走的路径消费
	visited   [110]int        //城市是否已经走过
	minL      [110][10010]int //minL[i][j]表示从1到i的所花费的j,而形成的最短路径长度
)

type Road struct {
	d, L, t int
	//这条路的d为终点,长度L,过路费
}

func dfs(s int) { //从s开始向N行走
	if s == N {
		minLen = utils.Min(minLen, totalLen)
		return
	}

	for i := 0; i < len(cityMap[s]); i++ {
		r := cityMap[s][i]
		cost := totalCost + r.t
		//r.d没走过才去走.
		if visited[r.d] == 0 {
			// 优化 1. 当前的花费超过总钱数,则不要dfs了,
			// 优化 2. 已经走过的路+这条路的长度 >= 最优长度,说明后续的路不用走了
			// 优化 3.原理: 如果达到某个状态A,发现前面也曾达到过状态A,且前面那次达到A的花销小,则剪枝
			// 		  并保存达到状态A的最少代价.
			// 优化 3. 走到城市k所花费的路费为m的条件下,最优长度为minL[k][m],若在后续
			//		  搜索中,再次走到k的时候,如果总路费恰好为m,且此时的长度L已经超过了
			//		  minL[k][m],则不必走了
			if  cost > K || totalLen+r.L >= minLen ||
				totalLen+r.L >= minL[r.d][cost]{
				continue
			}
			minL[r.d][totalCost+r.t] = totalLen + r.L
			totalLen += r.L
			totalCost += r.t
			visited[r.d] = 1
			dfs(r.d) //从r.d走完
			visited[r.d] = 0
			totalLen -= r.L
			totalCost -= r.t
		}
	}
}

func main() {
	fmt.Scanln(&K)
	fmt.Scanln(&N)
	fmt.Scanln(&R)
	for i := 0; i < R; i++ {
		var s int
		var r Road
		fmt.Scanln(&s, &r.d, &r.L, &r.t)
		if s != r.d {
			cityMap[s][i] = r
		}
	}

	totalLen, totalCost = 0, 0
	minLen = 1 << 30
	visited[1] = 1
	for i := 0; i < 110; i++ {
		for j := 0; j < 10010; j++ {
			minL[i][j] = 1 << 30
		}
	}
	dfs(1)
	if minLen < (1 << 30) {
		fmt.Println(minLen)
	} else {
		fmt.Println("none road")
	}
}


