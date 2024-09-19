package sort

func bubbleSort(arr []int) {

	length := len(arr)
	if length <= 0 {
		return
	}

	sorted := false
	n := length
	for ; !sorted ; {

		// 冒泡排序：每一轮循环，大泡泡冒泡到最后面
		sorted = true

		// arr[i] arr[i+1] => i: [0,n-1)
		for i := 0; i < n - 1; i++ {
			if arr[i] > arr[i + 1] {
				sorted = false
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		} 

		n -= 1

	}



}