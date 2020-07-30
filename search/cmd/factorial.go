/*
@Time : 2019/8/20 10:41
@Author : louis
@File : factorial
@Software: GoLand
*/

package main

import (
	"fmt"
	"github.com/oldthreefeng/algo/search"
)

func main() {
	a := search.Factorial(5)
	b := search.Fib(5)
	fmt.Println(a, b)
}
