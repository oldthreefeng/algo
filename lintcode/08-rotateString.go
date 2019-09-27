/*
@Time : 2019/9/26 14:45
@Author : louis
@File : 08-rotateString
@Software: GoLand
*/

package lintcode



func RotateString(str *string, offset int) {
	// write your code here
	l := len(*str)
	if str == nil || l == 0 {
		return
	}
	off := offset % l
	str1 := (*str)[:l-off]
	str2 := (*str)[l-off:]
	*str = str2 + str1
}
