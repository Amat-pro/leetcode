package sort


func quickSort(arr []int) {

	length := len(arr)
	if length <= 0 {
		return
	}

	left := 0
	right := length - 1

	_quickSort(arr, left, right)

}

func _quickSort(arr []int, left, right int) {

	// 确定终止条件
	if left >= right {
		return
	}

	middle := partition(arr, left, right)
	_quickSort(arr, left, middle-1)
	_quickSort(arr, middle+1, right)

}

func partition(arr []int, low int, high int) int {

	pivot := arr[low] // 选取第一个元素作为pivot

	for ; low < high ; {
		for ; low < high && arr[high] >= pivot ; {
			high--
		}
		arr[low] = arr[high]
		for ;  low < high && arr[low] <= pivot ; {
			low++
		}
		arr[high] = arr[low]
	}

	arr[high] = pivot
	return high

}

