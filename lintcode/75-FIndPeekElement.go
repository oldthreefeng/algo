/*
Copyright 2019 louis.
@Time : 2019/9/29 15:51
@Author : louis
@File : 75-FIndPeekElement
@Software: GoLand

*/

package lintcode

import "fmt"

func FindPeak(A []int) int  {
	for k,v := range A {
		if k ==0 || k == len(A)-1 {
			continue
		}
		fmt.Println(k, A[k-1],A[k], A[k+1])
		//fmt.Println(v)
		if  v >= A[k-1] && v >= A[k+1] {
			return k
		}
	}
	return  -1
}

