/*
Copyright 2019 louis.
@Time : 2019/9/28 22:01
@Author : louis
@File : 56-Twosum
@Software: GoLand

*/

package lintcode

//O(N)
func TwoSum(a []int, m int) []int {
	tmp := make(map[int]int)
	for k, v := range a {
		n := m - v               //找到要找的数n
		if x, ok := tmp[n]; ok { //在map找n,按索引查找,o(1)
			return []int{x, k}
		}
		tmp[v] = k //将a里面的索引和相应的m-v存入map
	}
	return nil
}

//O(N^2)
func TwoSumSlow(a []int, m int) []int {
	var res []int
	for i:=0;i<len(a); i++ {
		for j:=1;j<len(a);j++ {
			if a[i] + a[j] == m {
				return append(res,i,j)
			}
		}
	}
	return nil
}
