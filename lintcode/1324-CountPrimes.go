/*
Copyright 2019 louis.
@Time : 2019/9/26 23:11
@Author : louis
@File : 1324-CountPrimes
@Software: GoLand

*/

package lintcode

func IsPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func CountPrimes(n int) int {
	count := 0
	if n <= 2 {
		return count
	}
	for i := 2; i < n; i++ {
		if IsPrime(i) {
			count++
		}
	}
	return count
}

