/*
Copyright 2019 louis.
@Time : 2019/10/9 21:04
@Author : louis
@File : 125-isPalindrome
@Software: GoLand

*/

package leetcode

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	var str []rune = make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			str = append(str, unicode.ToLower(r))
		} else if unicode.IsDigit(r) {
			str = append(str, r)
		}
	}
	n := len(str) / 2
	for i := 0; i < n; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func isAlphaNum(c byte) bool {
	if c >= 'a' && c <='z' || c >= 'A' && c <= 'Z' ||
		c >= '0' && c <='9' {
		return true
	}

	return false
}


func isEqual(left, right byte) bool {
	if left >= 'a' && left <='z' && right >= 'a' && right <='z' ||
		left >= 'A' && left <='Z' && right >= 'A' && right <='Z' ||
		left >= '0' && left <='9' && right >= '0' && right <='9' {
		return left == right
	} else if left >= 'a' && left <='z' && right >= 'A' && right <='Z' {
		return left == (right + 32)
	} else if left >= 'A' && left <='Z' && right >= 'a' && right <='z' {
		return left == (right - 32)
	} else {
		return false
	}
}
func IsPalindromeX(s string) bool {

	i,j := 0, len(s)-1

	for i < j {
		if isAlphaNum(s[i]) == false {
			i++
			continue
		}
		if isAlphaNum(s[j]) == false {
			j--
			continue
		}

		if !isEqual(s[i], s[j]) {
			//fmt.Println(i, j, s[i], s[j])
			return false
		} else {
			i++
			j--
			continue
		}
	}
	return true
}