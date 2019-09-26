/*
@Time : 2019/9/26 14:06
@Author : louis
@File : sliceEqual
@Software: GoLand
*/

package utils


import "reflect"

func StringSliceReflectEqual(a,b []string) bool {
	return reflect.DeepEqual(a,b)
}

func StringSliceEqual(a,b []string) bool  {
	if len(a) != len(b) {return false}
	if (a==nil) != (b==nil) {return false}
	for i,v:= range a{
		if v!= b[i] {return false}
	}
	return true
}

func StringSliceEqualBce(a,b []string) bool  {
	if len(a) != len(b) {return false}
	if (a==nil) != (b==nil) {return false}
	b = b[:len(a)]
	for i,v:= range a{
		if v!= b[i] {return false}
	}
	return true
}

func IntSliceEqualBce(a,b []int) bool  {
	if len(a) != len(b) {return false}
	if (a==nil) != (b==nil) {return false}
	b = b[:len(a)]
	for i,v:= range a{
		if v!= b[i] {return false}
	}
	return true
}

