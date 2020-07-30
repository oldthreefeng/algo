/*
@Time : 2019/8/20 10:35
@Author : louis
@File : phoneFind
@Software: GoLand
*/

package main

import (
	"fmt"
	"github.com/oldthreefeng/algo/search"
)

func main() {
	pr, err := search.Find("15821586468")
	if err != nil {
		panic(err)
	}
	fmt.Println(pr)
}
