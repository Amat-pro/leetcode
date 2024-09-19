package sort

func selectionSort(arr []int) {

	length := len(arr)
	if length <= 0 {
		return
	}

	// 选择排序： 每一轮对于i,选择一个最小值j,交换位置
	var minIndex int
	for i := 0; i < length; i++ {
		minIndex = i
		for j := i+1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]

	}


}