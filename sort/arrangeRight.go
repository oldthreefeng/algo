/*
@Time : 2019/9/8 22:31
@Author : louis
@File : arrangeRight
@Software: GoLand
*/

package main

/*
输出前m大的数
用分治处理：复杂度 O(n+mlogm)
思路：把前m大的都弄到数组最右边，然后对这最右边m个元素排序，再输出
关键 ：O(n)时间内实现把前m大的都弄到数组最右边

引入操作 arrangeRight(k): 把数组(或数组的一部分）前k大的都弄到最右边
如何将前k大的都弄到最右边
1）设key=a[0], 将key挪到适当位置，使得比key小的元素都在key左边，比key大的元素都在key右边（线性时间完成）

2) 选择数组的前部或后部再进行 arrangeRight操作

a = k done 表示完成此过程
a > k 对右边a-1个元素再进行arrangeRigth(k)
a < k 对左边b个元素再进行arrangeRight(k-a)
*/

func arrangeRight(a []int,s,e,k int)  {
	if s >= e {
		return
	}
	t := a[s]
	i, j := s, e
	//前a[s]大的全部排到右边
	for i != j {
		for i < j && t <= a[j] {
			j--
		}
		for i < j && t >= a[i] {
			i++
		}
		Swap(a, i, j)
	}
	// 如果当前的i的位置等于k.前k大的已经排序玩必
	num := e - i + 1
	if num == k {
		return
	}
	if num > k {
		arrangeRight(a, i+1, e, k)
	}
	if num < k {
		arrangeRight(a, s, i-1, k-num)
	}
}

func main() {
	
}
