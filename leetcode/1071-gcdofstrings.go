/*
Copyright 2019 louis.
@Time : 2019/10/10 0:32
@Author : louis
@File : 1071-gcdofstrings
@Software: GoLand

*/

package leetcode

import (
	"github.com/oldthreefeng/algo/utils"
)

func GcdOfStrings(a, b string) string {
	if a+b != b+a {
		return ""
	}
	return a[0:utils.Gcd(len(a),len(b))]
}
