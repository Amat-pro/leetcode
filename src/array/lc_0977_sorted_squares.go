package array

// sortedSquares 有序数组的平方
func sortedSquares(nums []int) []int {

	if len(nums) == 0 {
		return nums
	}

	result := make([]int, len(nums))
	k := len(nums) - 1

	left := 0
	right := len(nums) - 1

	tempLeft := 0
	tempRight := 0

	for left <= right {

		tempLeft = nums[left] * nums[left]
		tempRight = nums[right] * nums[right]
		if tempLeft >= tempRight  {
			result[k] = tempLeft
			left++
		} else {
			result[k] = tempRight
			right--
		}

		k--

	}


	return result

}