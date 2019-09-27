/*
Copyright 2019 louis.
@Time : 2019/9/26 20:57
@Author : louis
@File : 01-repalce
@Software: GoLand

*/

package main

import (
	"fmt"
	"gogs.wangke.co/go/algo/lintcode"
)

func main() {
	a := 'e'
	b := 'd'
	str := "hello world"
	newStr := lintcode.StrReplaceCh(str, a, b)
	fmt.Println(newStr)

	ns := lintcode.StrDeleteCh(str, a)
	fmt.Println(ns)

	r1 := "abcdef"
	r2 := "bcd"  //index = 1的位置
	index := lintcode.StrStr(r1, r2)
	fmt.Println(index)
}
