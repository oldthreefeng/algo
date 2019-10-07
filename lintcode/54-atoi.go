/*
Copyright 2019 louis.
@Time : 2019/9/28 17:54
@Author : louis
@File : 54-atoi
@Software: GoLand

*/

package lintcode

import "strings"

const INT_MAX  = 1<<31 - 1
const INT_MIN = -(1 << 31)

func Atoi(s string) int {
	var tmp int
	var flags bool
	if s == "" {
		return 0
	}
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	for k,v := range s {
		if k== 0 {
			if v== '+' {
				continue
			} else if v == '-' {
				flags = true
				continue
			}
		}
		if v < '0' || v > '9' {
			break
		}
		if v >= '0' && v<= '9' {
			tmp = tmp*10 + int(v-'0')
		}
		if tmp > INT_MAX {
			break
		}
	}
	if flags {
		tmp = -tmp
	}
	if tmp > INT_MAX {
		return INT_MAX
	}
	if tmp < INT_MIN {
		return INT_MIN
	}
	return int(tmp)
}

func AtoiX(s string) int {
	var tmp int32
	var flags bool
	for k, v := range s {
		if v == ' ' {
			continue
		}
		if  v == '-' && tmp == 0 {
			flags = true
			continue
		}
		if v == '+' && tmp == 0  {
			continue
		}
		if tmp !=0 && (v<'0' || v>'9') {
			if flags {
				return -int(tmp)
			}
			return int(tmp)
		}
		if '0' <= v && v <= '9' {
			tmp = tmp*10 + v - '0'
		}
		if k > 31 {
			if flags {
				return INT_MIN
			}
			return INT_MAX
		}
	}
	if flags {
		return -int(tmp)
	}
	return int(tmp)
}
