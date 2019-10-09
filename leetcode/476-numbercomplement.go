/*
Copyright 2019 louis.
@Time : 2019/10/7 23:51
@Author : louis
@File : 476-numbercomplement
@Software: GoLand

*/

package leetcode

/*
FindComplement 生成某数的补数 `110(6) ==> 001(1)`
从右边开始,对第i位的处理,num已经右移了i次,此时对num的个位数位进行异或;将生成的异或位也左移i位.
然后将所有生成的异或位相加即可得到改数的`补数`.
*/
func FindComplement(num int) int {
	var c, tmp int
	var i uint
	for num != 0 {
		tmp = (num & 0x01) ^ 0x01//取num当前位的异或位
		c += tmp<<i  //将num当前位左移i位后
		num >>= 1
		i++
	}
	return c
}
