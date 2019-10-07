/*
Copyright 2019 louis.
@Time : 2019/10/7 23:51
@Author : louis
@File : 476-numbercomplement
@Software: GoLand

*/

package leetcode

func FindComplement(num int) int {
	var c, tmp int
	var i uint
	for num != 0 {
		tmp = (num & 0x01) ^ 0x01//取num当前位的异或
		c += tmp<<i  //将num当前位左移i位后
		num >>= 1
		i++
	}
	return c
}
