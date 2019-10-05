/*
Copyright 2019 louis.
@Time : 2019/9/30 20:36
@Author : louis
@File : 1095-findinMountingarray
@Software: GoLand

*/

package leetcode

type MountainArray struct {
	arr []int
}

func (m *MountainArray) get(index int) int {
	return m.arr[index]
}

func (m *MountainArray) length() int {
	return len(m.arr)
}

func bs(m *MountainArray, t, l, r int, asc bool) int {
	for l < r {
		mid := l + (r-l)>>1
		if (asc && m.get(mid) >= t) || (!asc && m.get(mid) <= t) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if t == m.get(l) {
		return l
	}
	return -1
}

func FindInMountainArray(target int, m *MountainArray) int {
	p, r := 0, m.length()-1
	for p < r {
		mid := p + (r-p)>>1
		// 如果mid值大于mid+1;说明峰值在mid的左边.否则只能在右边
		if m.get(mid) > m.get(mid+1) {
			r = mid
		} else {
			p = mid + 1
		}
	}
	//此次循环找到了p对应是峰值.左边进行查找,优先查找下标小的.
	i := bs(m, target, 0, p-1, true)
	if i != -1 {
		return i
	}
	//在右边进行查找
	return bs(m, target, p, m.length()-1, false)
}
