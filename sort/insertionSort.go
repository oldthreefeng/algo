package sort


func InsertionSort(arr []int) {
	lens := len(arr)
	for i := 1; i < lens; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			// 取出下一个元素，在已经排序的元素序列中从后向前扫描，
			// 如果该元素大于新元素，将该元素移到下一位置；把arr[j]和arr[j-1]交换,
			Swap(arr, j, j-1)
		}
	}
}

func UpInsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i] //从第二个元素开始比较.暂存arr[i]这个要插入的数据
		var j int
		for j = i - 1; j >= 0 &&
			//从小到大的顺序排序.
			arr[j] > insertVal; /*左边的数比要插入的数据大,继续找 */
		j-- {
			arr[j+1] = arr[j] //将arr[j]右移一位.
		}
		arr[j+1] = insertVal //如果j == i-1,没有进行比较,赋值依旧可以
	}
}

