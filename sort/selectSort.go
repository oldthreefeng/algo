package sort

func SelectSort(arr *[5]int) *[5]int {
	// *取值符 ， &取址符
	lens := len(arr)
	for i := 0; i < lens-1; i++ {
		minIndex := i
		for j := i + 1; j < lens; j++ {
			if arr[j] < arr[minIndex] { //找到最小的元素下标
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
		//tmp := arr[i]
		//arr[i] = arr[minIndex]
		//arr[minIndex] = tmp
	}
	return arr
}


//TODO  other Algorithm in sort
// add MergeSort 20190821
// add QuickSort 20190821
// add arrangeRight 20190908
