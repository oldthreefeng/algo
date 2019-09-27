/*
@Time : 2019/9/26 16:07
@Author : louis
@File : 13-strStr
@Software: GoLand
*/

package lintcode

func StrStr(source, target string) int {
	a, b := len(target), len(source)
	if a == 0 {
		return 0
	}
	if b == 0 || a > b {
		return -1
	}
	for i := 0; i < b-a+1; i++ {
		tmp := source[i:a+i]
		if tmp == target {
			return i
		}
	}
	return -1
}

func StrReplaceCh(str string, ch1, ch2 int32) string {
	tmp := []rune(str)
	for k, v := range tmp {
		if v == ch1 {
			tmp[k] = ch2
		}
	}
	return string(tmp)
}

func StrDeleteCh(str string, ch int32) string {
	tmp := []rune(str)
	for k, v := range tmp {
		if v == ch {
			tmp = append(tmp[:k], tmp[k+1:]...)
		}
	}
	return string(tmp)
}