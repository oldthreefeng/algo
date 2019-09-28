/*
Copyright 2019 louis.
@Time : 2019/9/28 23:14
@Author : louis
@File : 49-sortLetters
@Software: GoLand

*/

package lintcode

func SortLetters(st string) string {
	s := 0
	e := len(st)-1
	str := []rune(st)
	for s < e {
		for s <= e && str[s] > ']' {
			s++
		}
		for s<=e && str[e] <= ']' {
			e--
		}
		if s<=e {
			str[s],str[e] = str[e],str[s]
			s++
			e--
		}
	}
	return string(str)
}
