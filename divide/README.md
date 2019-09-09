# 分治应用

## 归并排序

```
数组排序任务可以如下完成：
1)  把前一半排序
2)  把后一半排序
3)  把两半归并到一个新的有序数组，然后再拷贝回原数组，排序完成。
```

[](../sort/MergeSort.go)

## 找数据前m大数据

```go

func ArrangeRight(arr []int, s, e, k int) {
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
	}else if k < a {
		ArrangeRight(arr, i+1, e, k)
	}else {
		ArrangeRight(arr, s, i-1, k-a)
	}

}

```