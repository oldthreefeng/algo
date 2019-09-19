/**
 * @time: 2019-08-20 22:49
 * @author: louis
 */
package search

import "sort"

/*
输入n ( n<= 100,000)个整数，找出其中的两个数，
它们之和等于整数m(假定肯定有解)。题中所有整数都能用int 表示 //百万量级的复杂度没问题.
*/

//解法1：用两重循环，枚举所有的取数方法，复杂度是O(n2)的。
//解法2:1)将数组排序，复杂度是O(n×log(n))
//     2)对数组中的每个元素a[i],在数组中二分查找m-a[i]，看能否找到。
//     复杂度log(n)，最坏要查找n-2次，所以查找这部分的复杂度也是O(n×log(n))这种解法总的复杂度是O(n×log(n))的。25

//解法3: 将数组遍历,存入map,key值为原数组索引,原数组索引存入value值,如果能从map里面取出要找的数,说明存在,复杂度为o(n)
func TwoSum(a []int, m int) []int {
	tmp := make(map[int]int)
	for k, v := range a {
		n := m - v               //找到要找的数n
		if x, ok := tmp[n]; ok { //在map找n,按索引查找,o(1)
			return []int{k, x}
		}
		tmp[v] = k //将a里面的索引和相应的m-v存入map
	}
	return nil
}

func TwoSunSort(a []int, m int) []int {
	sort.Ints(a)
	return nil
}
