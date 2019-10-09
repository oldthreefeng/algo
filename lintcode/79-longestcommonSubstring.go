/*
Copyright 2019 louis.
@Time : 2019/10/9 23:45
@Author : louis
@File : 79-longestcommonSubstring
@Software: GoLand

*/

package lintcode

import "gogs.wangke.co/go/algo/utils"

/*
dp[i][j]表示A串匹配到i，B串匹配到j时的最大公共长度
   dp[i][j]=dp[i-1][j-1]+1      ,A[i]==B[j]
   dp[i][j]=0                   ,A[i]!=B[j]
*/
func Lss(a, b string) int {
	var maxLen int
	m, n := len(a), len(b)
	if m < n {
		m, n = n, m
		a, b = b, a
	}
	trail := make([]int, n+1)
	for i:=0;i<m ;i++ {
		for j:=n-1;j>=0;j--{
			if a[i] == b[j] {
				trail[j+1] = 1+trail[j]
				maxLen = utils.Max(maxLen,trail[j+1])
			} else {
				trail[j+1] =0
			}
		}
	}
	return maxLen
}

func strLss(a, b string) string {
	var maxLen int
	l1 := len(a)
	l2 := len(b)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
			}
		}
	}
	var z int
	// s的长度必须是Lcs的长度
	var s []byte = make([]byte, maxLen, l1)
	for l1 != 0 && l2 != 0 {
		if a[l1-1] == b[l2-1] {
			l1--
			l2--
			s[z] = a[l1]
			z++
		} else if dp[l1-1][l2] < dp[l1][l2-1] {
			l2--
		} else if dp[l1][l2-1] <= dp[l1-1][l2] {
			l1--
		}
	}
	return string(s)
}
