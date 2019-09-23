/*
@Time : 2019/9/21 22:01
@Author : louis
@File : getinverse
@Software: GoLand
*/

package main

import (
	"fmt"
	"gogs.wangke.co/go/algo/sort"
)

var inv [100]int

func getInverse(t, p int) int {
	if t == 1 {
		return 1
	} else {
		return (p - p/t) * getInverse(p%t, p) % p
	}
}

func main() {
	fmt.Println(getInverse(17, 26))


	fmt.Println(sort.ExGcd(17, 26))
}
