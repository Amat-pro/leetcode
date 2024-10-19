package temp

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(x, y int) bool {
		if nums[x] <= nums[y] {
			return true
		}
		return false
	})

	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		// 对i去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		var temp int
		for left < right {
			temp = nums[i] + nums[left] + nums[right]
			if temp > 0 {
				right--
			} else if temp < 0 {
				left++
			} else {
				// 收集结果
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// 对left去重
				// 对right去重
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}

	}

	return result
}
