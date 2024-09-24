package two_point

// fourSum 四数之和
func fourSum(nums []int, target int) [][]int {

	result := [][]int{}
	length := len(nums)

	selectSort(nums)

	for i := 0; i < length; i++ {

		// 对i去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < length; j++ {
			// 对j去重
			if nums[j] == nums[j-1] {
				continue
			}

			if nums[i]+nums[j] > target {
				return result
			}

			left := j + 1
			right := length - 1
			for left < right {
				temp := nums[i] + nums[j] + nums[left] + nums[right]
				if temp < target {
					left++
				} else if temp > target {
					right--
				} else {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
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

	}

	return result
}
