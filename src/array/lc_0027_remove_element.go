package array

// removeElements 移除元素
func removeElements(nums []int, val int) int {

	next_to_write_idx := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[next_to_write_idx] = nums[i]
			next_to_write_idx++
		}
	}

	return next_to_write_idx

}