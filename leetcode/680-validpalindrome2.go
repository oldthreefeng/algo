/*
@Time : 2019/10/9 11:52
@Author : louis
@File : 680-validpalindrome2
@Software: GoLand
*/

package leetcode

// ValidPalindrome 判断一个str是否为
func ValidPalindrome(s string) bool {
	//	return len(s) - Lcs(s,utils.Reverse(s)) <=1 //使用动态规划，超出内存限制
	n := len(s)
	i, j := 0, n-1
	//左右检测第一个不同的字符
	for i <= j && s[i] == s[j] {
		i++
		j--
	}
	// s本身为回文字符串序列
	if i > j {
		return true
	}
	// 得到形如a****b的子串; 判断a****或者****b是否为回文序列
	return palindrome(s[i:j]) || palindrome(s[i+1:j+1])
}

func palindrome(s string) bool {
	n := len(s)
	i, j := 0, n-1
	for i <= j && s[i] == s[j] {
		i++
		j--
	}
	if i > j {
		return true
	} else {
		return false
	}
}
