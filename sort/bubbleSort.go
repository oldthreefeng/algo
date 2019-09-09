package sort

func BubbleSort(arr []int) {

	lens := len(arr)
	tmp := 0
	for j := 0; j < lens-1; j++ {
		for i := 0; i < lens-1-j; i++ {
			if arr[i] > arr[i+1] {
				tmp = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
			}
		}
	}


}


