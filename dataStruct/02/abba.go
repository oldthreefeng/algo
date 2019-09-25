/*
Copyright 2019 louis.
@Time : 2019/9/25 21:49
@Author : louis
@File : abba
@Software: GoLand

*/

package main

import "fmt"

func reverse(s string) string  {
	b := []rune(s)
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func abba(s string) bool  {
	return s == reverse(s)
}

func main() {
	s := "abbab"
	if abba(s) {
		fmt.Print("yes")
	}
}