package search

//LowerBound 在从小到大的数组里面
func LowerBound(a []int, size, p int) int {
	L := 0	//左端点
	R := size - 1 //右端点
	lastPos := -1 //到目前为止最右的为止
	for L < R {
		mid := L + (R-L)/2 //求区间中点, 防止L+R过大溢出
		if p >= a[mid] {
			R = mid -1
		} else {
			lastPos = mid
			L = mid +1
		}
	}
	return lastPos
}
