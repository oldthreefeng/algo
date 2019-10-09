/*
Copyright 2019 louis.
@Time : 2019/10/8 0:31
@Author : louis
@File : 231-ispoweroftwo
@Software: GoLand

*/

package leetcode

//IsPowerOfTwo
//n为2的幂，二进制表示时：一定为 某位是1，其它位都是0，例如4：00000100
//n-1,二进制表示时，n的二进制为1的位右边全是1，例如3：00000011
//所以 & 运算后，结果是 00000000, 则必定为 true
func IsPowerOfTwo(n int) bool {
	if n <=0 {
		return false
	}

	return (n &(n-1)) == 0
}
