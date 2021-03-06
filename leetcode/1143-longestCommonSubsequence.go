/*
@Time : 2019/10/8 14:32
@Author : louis
@File : 1143-longestCommonSubsequence
@Software: GoLand
*/

package leetcode

import "github.com/oldthreefeng/algo/utils"

/*
Lcs 获取string的最长公共子序列

- 思路

```cgo
输入两个串s1,s2,
    设MaxLen(i,j)表示:  s1的左边i个字符形成的子串，与s2左边的j个字符形成的子串的最长公共子序列的长度(i,j从0开始算）
    MaxLen(i,j) 就是本题的“状态”
    假定len1 = strlen(s1),len2 = strlen(s2）
    那么题目就是要求MaxLen(len1,len2)
    显然：
    MaxLen(n,0)  = 0  ( n= 0...len1）
    MaxLen(0,n)  = 0  ( n=0...len2）递推公式：
    if ( s1[i-1] == s2[j-1] ) //s1的最左边字符是s1[0]
        MaxLen(i,j) = MaxLen(i-1,j-1) + 1
    else
        MaxLen(i,j) = Max(MaxLen(i,j-1),MaxLen(i-1,j) )
```

- 证明

```cgo
S1,S2表示表示长度为i,j的序列
S1[:i-1]表示长度为i-1的序列
S2[:j-1]表示长度为j-1的序列
S1[i-1]表示S1的第i-1个字符
S2[j-1]表示S2的第j-1个字符

在S1[i-1] != S2[j-1]情况下:
MaxLen(S1,S2)不会比MaxLen(S1,S2[:j-1])和MaxLen(S1[:i-1],S2)两者之中任意一个小,也不会比两者都大.
	证明:
采用反证法:
    若MaxLen(S1,S2) 比 MaxLen(S1,S2[:j-1])和MaxLen(S1[:i-1],S2) 都大,
    情况1: MaxLen(S1,S2) > MaxLen(S1,S2[:j-1]
    则,S1[i-1] 必定是 MaxLen(S1,S2)的子序列的关键元素,且为最后一个元素.
    同理:S2[j-1] 必定是 MaxLen(S1,S2)的子序列的关键元素,且为最后一个元素.
    故而S1[i-1]==S2[j-1],
    与假设矛盾,从而假设不成立. 证毕~

```
*/
func Lcs(text1 string, text2 string) int {
	l1 := len(text1)
	l2 := len(text2)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = utils.Max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[l1][l2]
}

func StrLcs(text1 string, text2 string) string {
	l1 := len(text1)
	l2 := len(text2)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = utils.Max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	var z int
	// s的长度必须是Lcs的长度
	var s []byte = make([]byte,dp[l1][l2], l1)
	for l1 != 0 && l2 != 0 {
		if text1[l1-1] == text2[l2-1] {
			l1--
			l2--
			s[z] = text1[l1]
			z++
		} else if dp[l1-1][l2] < dp[l1][l2-1] {
			l2--
		} else if dp[l1][l2-1] <= dp[l1-1][l2] {
			l1--
		}
	}
	return string(s)
}
