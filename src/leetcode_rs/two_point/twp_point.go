package two_point

import "sort"

// threeSum 15 - 三数之和
func threeSum(nums []int) [][]int {

	result := make([][]int, 0, 10)

	// 排序
	sort.Slice(nums, func(x, y int) bool {
		return nums[x] < nums[y]
	})

	for i := 0; i < len(nums); i++ {

		// 对i去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 双指针
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

// fourSum 18 - 四数之和
func fourSum(nums []int, target int) [][]int {

	result := make([][]int, 0, 10)

	// 排序
	sort.Slice(nums, func(x, y int) bool {
		return nums[x] < nums[y]
	})

	for i := 0; i < len(nums); i++ {

		// 对i去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			// 对j去重
			if nums[j] == nums[j-1] {
				continue
			}

			sum := target - nums[i] - nums[j]

			left := j + 1
			right := len(nums) - 1
			// 双指针
			for left < right {
				temp := nums[left] + nums[right]
				if temp < sum {
					left++
				} else if temp > sum {
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
