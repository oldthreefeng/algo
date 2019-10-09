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

// FindInMountainArray  在山脉数组中找某数.
// 定义左右边界l和r，mid=(l+r)>>1，判断arr[mid]>arr[mid+1]是不是成立，如果成立，那么山峰在[0,mid]这个区间中，我们在这个区间中继续查找；
// 如果不成立，那么山峰在[mid+1,len(arr))这个区间中，我们在这个区间中查找。
// 最后山峰左边查找目标和山峰右边查找目标使用的就是二分查找，不过山峰左边是递增，山峰右边是递减.
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
