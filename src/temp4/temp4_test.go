package main

import "math"

// search - 二分查找
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	return _search(nums, 0, len(nums)-1, target)
}

func _search(nums []int, left, right int, target int) int {

	// stop
	if left == right {
		if nums[left] == target {
			return left
		}

		return -1
	}

	if left > right {
		return -1
	}

	// 递归
	middle := (left + right) / 2
	if nums[middle] > target {
		// 向左
		return _search(nums, left, middle-1, target)
	} else if nums[middle] < target {
		// 向右
		return _search(nums, middle+1, right, target)
	} else {
		return middle
	}

}

// minSubArrayLen - 长度最小的子数组
func minSubArrayLen(target int, nums []int) int {

	sum := 0
	minLength := math.MaxInt

	// [start,end]
	start := 0
	for end := 0; end < len(nums); end++ {
		sum += nums[end]

		if sum == target {
			length := end - start
			if length < minLength {
				minLength = length
			}
		}

		if sum > target {
			sum -= nums[start]
			start++
		}
	}

	if minLength == math.MaxInt {
		return -1
	}

	return minLength

}

type Node struct {
	Value int
	Next  *Node
}

// reverseList - 翻转链表
func reverseList(head *Node) *Node {

	pre := head
	cur := head.Next

	pre.Next = nil

	for cur != nil {
		next := cur.Next
		
		cur.Next = pre

		pre = cur
		cur = next

	}

	return pre

}

// 滑动窗口最大值 
// 每日温度
// 下一个更大元素
// 下一个更大元素II