/*
@Time : 2019/10/10 17:07
@Author : louis
@File : 20191010-1
@Software: GoLand
*/

package leetcode

func Fraction(cont []int) []int {
	var a, b, c int
	c = 1
	a = cont[len(cont)-1]
	for i := len(cont) - 2; i >= 0; i-- {
		b = cont[i]
		if b == 0 {
			a, c = c, a
		} else {
			var tmp = a
			a = a*b + c
			c = tmp
			helper(&a, &c)
		}
	}
	helper(&a, &c)
	var res []int
	res = append(res, a, c)
	return res
}

func helper(a, b *int) {
	g := gcd(*a, *b)
	*a /= g
	*b /= g
}

func gcd(a, b int) int {
	if a < 0 || b < 0 {
		return -1
	}
	if a%b != 0 {
		return gcd(b, a%b)
	}
	return b
}


func FractionX(cont []int) []int {
	v1, v2 := f(cont)
	return []int{v2, v1}
}

func f(cont []int) (int, int) {
	if len(cont) == 1 {
		return 1, cont[0]
	}
	v1, v2 := f(cont[1:])
	return  v2, cont[0] * v2 + v1
}