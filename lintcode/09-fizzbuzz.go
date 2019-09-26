/*
@Time : 2019/9/26 15:08
@Author : louis
@File : 09-fizzbuzz
@Software: GoLand
*/

package lintcode

import "fmt"

/*
给你一个整数n. 从 1 到 n 按照下面的规则打印每个数：

如果这个数被3整除，打印fizz.
如果这个数被5整除，打印buzz.
如果这个数能同时被3和5整除，打印fizz buzz.
如果这个数既不能被 3 整除也不能被 5 整除，打印数字本身。
Example
比如 n = 15, 返回一个字符串数组：

[
  "1", "2", "fizz",
  "4", "buzz", "fizz",
  "7", "8", "fizz",
  "buzz", "11", "fizz",
  "13", "14", "fizz buzz"
]
Challenge
你是否可以只用一个 if 来实现
*/
func FizzBuzz(n int) []string {
	// write your code here
	a := make([]string, 0)
	for i := 0; i < n; i++ {
		switch {
		case i%15 == 0:
			a = append(a, "fizz buzz")
		case i%3 == 0:
			a = append(a, "fizz")
		case i%5 == 0:
			a = append(a, "buzz")
		default:
			a = append(a, fmt.Sprint(i))
		}
	}
	return a
}
