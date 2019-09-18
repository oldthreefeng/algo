/*
@Time : 2019/9/18 1:21
@Author : louis
@File : maxlentwostring
@Software: GoLand
*/

package main

import (
	"fmt"
	"gogs.wangke.co/go/algo/utils"
)

var (
	maxLen   [1000][1000]int
	sz1, sz2 string
)

func main() {
	_, _ = fmt.Scanln(&sz1, &sz2)
	length1 := len(sz1)
	length2 := len(sz2)
	for i := 0; i <= length1; i++ {
		maxLen[i][0] = 0
	}
	for i := 0; i <= length2; i++ {
		maxLen[0][i] = 0
	}

	for i := 1; i <= length1; i++ {
		for j := 1; j <= length2; j++ {
			if sz1[i-1] == sz2[j-1] {
				maxLen[i][j] = maxLen[i-1][j-1] + 1
			} else {
				maxLen[i][j] = utils.Max(maxLen[i][j-1], maxLen[i-1][j])
			}
		}
	}
	fmt.Println(maxLen[length1][length2])
}
