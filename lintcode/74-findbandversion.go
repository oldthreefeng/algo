/*
Copyright 2019 louis.
@Time : 2019/9/29 15:28
@Author : louis
@File : 74-findbandversion
@Software: GoLand

*/

package lintcode

func FindBadVersion(n int) int {
	for i := 1; i <= n; i++ {
		if IsBadVersion(i) {
			return i
		}
	}
	return 0
}

func IsBadVersion(n int) bool {
	if n > 32 {
		return false
	}
	return true
}
