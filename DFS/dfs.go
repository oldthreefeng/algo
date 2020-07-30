/*
@Time : 2019/9/18 23:53
@Author : louis
@File : dfs
@Software: GoLand
*/

package main

import (
	"fmt"
	"github.com/oldthreefeng/algo/utils"
)

var (

	rooms, color [60][60]int
	maxRooArea   = 0
	roomNum      = 0
	roomArea     int
)

/*
输入
程序从标准输入设备读入数据。
第一行是两个整数，分别是南北向、东西向的方块数。
在接下来的输入行里，每个方块用一个数字(0≤p≤50)描述。
用一个数字表示方块周围的墙，1表示西墙，2表示北墙，4表示东墙，8表示南墙。
每个方块用代表其周围墙的数字之和表示。
城堡的内墙被计算两次，方块(1,1)的南墙同时也是方块(2,1)的北墙。
输入的数据保证城堡至少有两个房间。

输出
城堡的房间数、城堡中最大房间所包括的方块数。
结果显示在标准输出设备上。


描述

```
     1   2   3   4   5   6   7
   #############################
 1 #   |   #   |   #   |   |   #
   #####---#####---#---#####---#
 2 #   #   |   #   #   #   #   #
   #---#####---#####---#####---#
 3 #   |   |   #   #   #   #   #
   #---#########---#####---#---#
 4 #   #   |   |   |   |   #   #
   #############################
           (图 1)

   #  = Wall
   |  = No wall
   -  = No wall
```

图1是一个城堡的地形图。请你编写一个程序，计算城堡一共有多少房间，最大的房间有多大。城堡被分割成mn(m≤50，n≤50)个方块，每个方块可以有0~4面墙。

输入

程序从标准输入设备读入数据。第一行是两个整数，分别是南北向、东西向的方块数。在接下来的输入行里，每个方块用一个数字(0≤p≤50)描述。用一个数字表示方块周围的墙，1表示西墙，2表示北墙，4表示东墙，8表示南墙。每个方块用代表其周围墙的数字之和表示。城堡的内墙被计算两次，方块(1,1)的南墙同时也是方块(2,1)的北墙。输入的数据保证城堡至少有两个房间。

输出

城堡的房间数、城堡中最大房间所包括的方块数。结果显示在标准输出设备上。

样例输入

```
4
7
11 6 11 6 3 10 6
7 9 6 13 5 15 5
1 10 12 7 13 7 5
13 11 10 8 10 12 13
```

样例输出

```
5
9
```

*/

//TODO fix range -1

func main() {
	var R, C int
	fmt.Scanln(&R)
	fmt.Scanln(&C)
	for i := 1; i <= R; i++ {
		for k := 1; k <= C; k++ {
			fmt.Scanln(rooms[i][k])
		}
	}
	for i := 1; i <= R; i++ {
		for k := 1; k <= C; k++ {
			if color[i][k] == 0 {
				roomNum++ //房间数加一
				roomArea = 0
				Dfs(i, k)
				maxRooArea = utils.Max(roomArea, maxRooArea)
			}
		}
	}
	fmt.Println(roomNum)
	fmt.Println(maxRooArea)
}

func Dfs(i, k int) {
	if color[i][k] != 0 {
		return //走过了
	}
	roomNum++
	color[i][k] = roomNum
	if rooms[i][k] & 1 == 0 { //西
		Dfs(i,k-1)
	}
	if rooms[i][k] & 2 == 0 { //北
		Dfs(i-1,k)
	}
	if rooms[i][k] & 4 == 0 { //东
		Dfs(i,k+1)
	}
	if rooms[i][k] & 8 == 0 { //南
		Dfs(i+1,k)
	}
}
