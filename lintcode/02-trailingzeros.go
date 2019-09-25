/*
@Time : 2019/9/25 16:25
@Author : louis
@File : 02
@Software: GoLand
*/

package lintcode

/*
设计一个算法，计算出n阶乘中尾部零的个数

Example
样例  1:
	输入: 11
	输出: 2

	样例解释:
	11! = 39916800, 结尾的0有2个。

样例 2:
	输入:  5
	输出: 1

	样例解释:
	5! = 120， 结尾的0有1个。

Challenge
O(logN)的时间复杂度
*/

//超时
func TrailingZeros(n int64) int64 {
	// write your code here, try to do it without arithmetic operators.
	// 先计算阶乘
	nFact := fact(n)
	// 在计算trailing.
	return counterFactor(nFact,10)
}

func counterFactor(m, factor int64) (res int64) {
	for m%factor == 0 {
		m = m/10
		res++
	}
	return res
}

func fact(n int64) (res int64) {
	res = 1
	var i int64
	for i = 1; i <= n; i++ {
		res *= i
	}
	return res
}

/*
原理:
$$
f(n)=[\farc{n!}{5}] + [\farc{n!}{5^2}] +\cdots + [\farc{n!}{5^n}]
$$
$$
q_0 = n
q_{i+1} = [\farc{q_i}{5}]
$$
$$
q_{i+1} = 0  <==> 5^{k+1} > n
$$

for q != 0 {
	q /= 5   // ==> a1 + a2 + ... + an
	res += q
}

*/
func TrailingZerosX(n int64) int64 {
	// write your code here, try to do it without arithmetic operators.
	var res int64

	for n > 0 {
		n /= 5
		res +=n
	}
	return res
}
