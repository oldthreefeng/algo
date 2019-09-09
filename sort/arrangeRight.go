/*
@Time : 2019/9/8 22:31
@Author : louis
@File : arrangeRight
@Software: GoLand
*/

package sort

// ArrangeRight
//输出前m大的数.
//用分治处理：复杂度 O(n+mlogm), 直接用MergeSort复杂度为O(n*logn).
//思路：把前m大的都弄到数组最右边，然后对这最右边m个元素排序，再输出.
//关键 ：O(n)时间内实现把前m大的都弄到数组最右边.
//
//引入操作 ArrangeRight(k): 把数组(或数组的一部分）前k大的都弄到最右边
//如何将前k大的都弄到最右边.
//
//1）设key=a[0], 将key挪到适当位置，使得比key小的元素都在key左边，
// 比key大的元素都在key右边（线性时间完成）
//
//2) 选择数组的前部或后部再进行 ArrangeRight操作
//a == k ,done 表示完成此过程,
//a > k 对右边a-1个元素再进行ArrangeRight(k),
//a < k 对左边b个元素再进行ArrangeRight(k-a).
func ArrangeRight(arr []int,s,e,k int)  {
	if s >= e {
		return
	}
	if k == e-s+1 {
		return
	}
	key := arr[s]
	i, j := s, e
	//前a[s]大的全部排到右边
	for i != j {
		for j > i && arr[j] >= key {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
		for j > i && arr[i] <= key {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	// 如果当前的t[s]的位置等于k.前k大的已经排序完毕
	a := e - i + 1
	if k == a {
		return
	//a > k 对右边a-1个元素再进行arrangeRight(k)
	}else if k < a {
		ArrangeRight(arr, i+1, e, k)
	}else {
		ArrangeRight(arr, s, i-1, k-a)
	}
}
