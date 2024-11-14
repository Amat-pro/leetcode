package array

import "math"

// search 0704 - 二分查找
// nums是有序的
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	return _search(nums, 0, len(nums)-1, target)

}

// _search 递归实现
func _search(nums []int, left, right int, target int) int {
	// 确定终止条件
	if left == right {
		if nums[left] == target {
			return left
		}
		return -1
	}
	if left > right {
		return -1
	}

	// [left,right]查找
	middle := (left + right) / 2
	if nums[middle] > target {
		return _search(nums, left, middle-1, target)
	} else if nums[middle] < target {
		return _search(nums, middle+1, right, target)
	} else {
		return middle
	}

}

// removeElement 27 - 移除元素
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	ptr1 := 0
	ptr2 := 0

	for ptr2 < len(nums) {
		if nums[ptr2] != val {
			nums[ptr1] = nums[ptr2]
			ptr1++
			ptr2++
			continue
		}
		ptr2++
	}

	return ptr1

}

// sortedSquares 977 - 有序数组的平方
func sortedSquares(nums []int) []int {

	// 双指针方法
	// 最大值一定出现在数组两边的位置
	// 用left,right动态确定剩余元素的边界[left,right]
	// k表示每次要写入的最大值的位置
	left := 0
	right := len(nums) - 1
	k := right

	result := make([]int, len(nums))
	for left <= right {
		leftTemp := nums[left] * nums[left]
		rightTemp := nums[right] * nums[right]
		if rightTemp >= leftTemp {
			result[k] = rightTemp
			right--
		} else {
			result[k] = leftTemp
			left++
		}

		k--
	}

	return result

}

// minSubArrayLen 209 - 长度最小的子数组
func minSubArrayLen(target int, nums []int) int {

	// 滑动窗口解法 O(n)
	sum := 0
	minLength := math.MaxInt

	start := 0
	for end := 0; end < len(nums); end++ { // 滑动窗口位置-end
		sum += nums[end]

		// 滑动窗口位置-start
		for sum > target {
			sum -= nums[start]
			start++
		}

		if sum == target {
			temp := end - start + 1
			if temp < minLength {
				minLength = temp
			}
		}

	}

	if minLength == math.MaxInt {
		return -1
	}

	return minLength

}

