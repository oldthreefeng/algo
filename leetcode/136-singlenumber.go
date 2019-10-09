/*
@Time : 2019/10/9 16:34
@Author : louis
@File : 136-singlenumber
@Software: GoLand
*/

package leetcode

/*
SingleNumber

交换律：`a ^ b ^ c <=> a ^ c ^ b`

任何数于0异或为任何数: `0 ^ n => n`

相同的数异或为0: `n ^ n => 0`
*/
func SingleNumber(nums []int) int {
	x := nums[0]
	for i := 0; i < len(nums); i++ {
		x ^= nums[i]
	}
	return x
}

/*
SingleNumber2
已知:
0 ^ x = x,
x ^ x = 0;
x & ^x = 0,
x & ^0 = x;
则有:
x出现一次:
	a = (a ^ x) & ^b ==>  a = x
	b = (b ^ x) & ^a ==> (因为a=x,所有b=0)
x出现两次:
	a = (a ^ x) & ^b ==>  a = (x ^ x) & ^0 ==> a = 0
	b = (b ^ x) & ^a ==>  b = (0 ^ x) & ^0 ==> b = x
x出现三次:
	a = (a ^ x) & ^b ==>  a = (0 ^ x) & ^x ==> a = 0
	b = (b ^ x) & ^a ==>  b = (x ^ x) & ^0 ==> b = 0
*/
func SingleNumber2(nums []int) int {
	a, b := 0, 0
	for _, x := range nums {
		a = (a ^ x) & ^b
		b = (b ^ x) & ^a
	}
	return a
}

func SingleNumber2x(nums []int) int {
	var (
		map1 = make(map[int]int8)
	)
	for _, value := range nums {
		map1[value]++
	}
	for key, v := range map1 {
		if v == 1 {
			return key
		}

	}
	return 0
}

/*
SingleNumber3
位运算，异或运算。对于一个数组`nums = [1, 1 , 2, 2, 3, 4, 4, 5]`。
其一，如果，进行一次全部异或运算，将会得到`3 xor 5`。
其二， `3 ^ 5 = 110b`。那么在出现`1`的位置，必然一个为`1`一个为`0`，这样就可以根据特征区分出两个数字。
其三，于是将问题转化为了“一个数字出现1次，其他数字出现两次”。
*/
func SingleNumber3(nums []int) []int {
	var diff int
	var res = make([]int, 2)
	for _, v := range nums {
		diff ^= v
	}
	//  3(011),5(101) 两个不一样，  diff = 110
	diff &= ^diff +1 //==> 找到只出现一次的两个数最右侧不相同的位
	// diff = 10
	for _, v := range nums {
		if v&diff == 0 {
			res[0] ^= v
		} else {
			res[1] ^= v
		}
	}
	return res
}
