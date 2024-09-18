package lc_0704_search

// search 左闭右闭
func search(numbers []int, target int) int {
	if len(numbers) == 0 {
		return -1
	}

	left := 0
	right := len(numbers)
	var middle int 

	for left <= right {
		middle = (left + right) / 2
		if target < numbers[middle] {
			right = middle - 1
		} else if target > numbers[middle] {
			left = middle + 1
		} else {
			return middle
		}
	
	}

	return -1

}