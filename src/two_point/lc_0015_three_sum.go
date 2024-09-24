package two_point

// threeSum 三数之和
func threeSum(nums []int) [][]int {

	if len(nums) < 3 {
		return nil
	}

	bubbleSort(nums) // 排序

	result := [][]int{}
	for i := 0; i < len(nums)-2; i++ {

		// 对i去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			temp := nums[i] + nums[left] + nums[right]
			if temp > 0 {
				right--
			} else if temp < 0 {
				left++
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				// 对left去重
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				// 对right去重
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return result
}

func selectSort(nums []int) {
	// 选择排序
	if len(nums) == 0 {
		return
	}

	minValIdx := 0
	for i := 0; i < len(nums)-1; i++ {
		minValIdx = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minValIdx] {
				minValIdx = j
			}
		}
		nums[i], nums[minValIdx] = nums[minValIdx], nums[i]
	}
}

func bubbleSort(nums []int) {
	if len(nums) == 0 {
		return
	}

	sorted := false
	n := len(nums)

	for !sorted {

		sorted = true

		for i := 0; i < n-1; i++ {
			if nums[i] > nums[i+1] {
				sorted = false
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}

		n -= 1
	}

}
