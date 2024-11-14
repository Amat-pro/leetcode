package queue_stack

import (
	"fmt"
	"sort"
	"strconv"
)

// isValid 20 - 有效的括号
func isValid(s string) bool {
	runes := []rune(s)
	stack := make([]rune, 0, 10)

	for i := 0; i < len(runes); i++ {
		switch runes[i] {
		case '(':
			stack = append(stack, runes[i])
		case '{':
			stack = append(stack, runes[i])
		case '[':
			stack = append(stack, runes[i])
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		default:
			return false
		}
	}

	return len(stack) == 0

}

// removeDuplicate 1047 - 删除字符串中的所有相邻重复项
func removeDuplicates(s string) string {

	runes := []rune(s)

	stack := make([]rune, 0, 10)
	for i := 0; i < len(runes); i++ {

		// 弹出
		if len(stack) > 0 && stack[len(stack)-1] == runes[i] {
			stack = stack[:len(stack)-1]
			continue
		}

		// 加入
		stack = append(stack, runes[i])

	}

	return string(stack)

}

// evalRPN 150 - 逆波兰表达式求值
func evalRPN(tokens []string) int {

	stack := make([]int, 0, 10)
	for _, token := range tokens {

		result := 0

		if token == "+" {
			i := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result = i + j

		} else if token == "-" {
			i := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result = j - i
		} else if token == "*" {
			i := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result = i * j

		} else if token == "/" {
			i := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result = int(j / i)
		} else {
			result, _ = strconv.Atoi(token)

		}

		stack = append(stack, result)

	}

	return stack[len(stack)-1]
}

// maxSlidingWindow 239 - 滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	queue := NewQueue()
	for i := 0; i < k; i++ {
		queue.push(nums[i])
	}

	result := make([]int, 0, 10)
	// [0, k-1]窗口最大值
	result = append(result, queue.getTop())

	popIndex := 0
	pushIndex := k
	for pushIndex <= len(nums)-1 {
		queue.push(nums[pushIndex])
		queue.pop(nums[popIndex])

		result = append(result, queue.getTop())

		popIndex++
		pushIndex++

	}

	return result

}

// 实现一个Queue - 单调栈
type Queue struct {
	q []int
}

func NewQueue() *Queue {
	queue := make([]int, 0, 10)
	return &Queue{
		q: queue,
	}
}

func (queue *Queue) push(num int) {

	if queue.len() == 0 {
		queue.q = append(queue.q, num)
		return
	}

	// 向前卷走每一个比num小的元素，因为这些元素不会成为窗口内最大值
	for queue.len() > 0 && num > queue.getBack() {
		queue.q = queue.q[:queue.len()-1]
	}

	queue.q = append(queue.q, num)

}

func (queue *Queue) pop(num int) {
	if queue.len() == 0 {
		return
	}

	if num != queue.getTop() {
		return
	}

	// 弹出的是最大值元素
	queue.q = queue.q[1:]

}

// getTop 等价于 getMaxValue
func (queue *Queue) getTop() int {
	return queue.q[0]
}

func (queue *Queue) getBack() int {
	return queue.q[queue.len()-1]
}

func (queue *Queue) len() int {
	return len(queue.q)
}

// topFrequent 347 - 前k个高频元素 - O(n*logn)
func topFrequent(nums []int, k int) []int {
	frequents := map[int]int{}

	for _, v := range nums {
		_, ok := frequents[v]
		if ok {
			frequents[v]++
		} else {
			frequents[v] = 1
		}
	}

	keys := make([]int, 0, len(frequents))
	for k := range frequents {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(x, y int) bool {
		return frequents[keys[x]] > frequents[keys[y]]
	})

	return keys[:k]

}

// ================================================== 单调栈 ==================================================

// dailyTemperatures 739 - 每日温度
func dailyTemperatures(temperatures []int) []int {

	stack := make([]int, 0, len(temperatures))

	result := make([]int, len(temperatures))
	for i, temperature := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}

		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)

	}

	return result

}

// nextGreaterElement 496 - 下一个更大元素I
func nextGreaterElement(nums1 []int, nums2 []int) []int {

	nums2Map := map[int]int{}
	for i, v := range nums2 {
		nums2Map[v] = i
	}

	nextGreater := make([]int, len(nums2)) // 存放下一个更大元素的位置index
	stack := make([]int, 0, len(nums2))
	for i, num := range nums2 {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}

		for len(stack) > 0 && num > nums2[stack[len(stack)-1]] {
			nextGreater[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	result := make([]int, len(nums1))
	for i, v := range nums1 {
		if nextGreater[nums2Map[v]] == 0 {
			result[i] = -1
		} else {
			result[i] = nums2[nextGreater[nums2Map[v]]]
		}
	}

	return result

}

// nextGreaterElementsII 503 - 下一个更大元素II
func nextGreaterElementsII(nums []int) []int {

	// 循环数组
	// index := virtualIndex % len(nums)
	l := len(nums) * 2
	stack := make([]int, 0, l)

	result := make([]int, len(nums)) // 存放下一个更大元素(因为是循环数组，存储下标时出现0的情况无法区分下一个更大元素是下标0还是初始值0)
	for i := range result {
		result[i] = -1
	}

	for i := 0; i < l; i++ {
		index := i % len(nums) // 重新计算index
		if len(stack) == 0 {
			stack = append(stack, index)
			continue
		}

		for len(stack) > 0 && nums[index] > nums[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = nums[index]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, index)
	}

	return result

}

// trap 42 - 接雨水
// 找左右两边第一个比stack[len[stack]-1]高的柱子
func trap(heights []int) int {

	stack := make([]int, 0, 10)

	result := 0
	for i, height := range heights {

		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}

		for len(stack) > 0 && height > heights[stack[len(stack)-1]] {
			middle := heights[stack[len(stack)-1]] // 中间柱子
			stack = stack[:len(stack)-1]

			// 找不到左边柱子的情况
			if len(stack) == 0 {
				break
			}

			left := heights[stack[len(stack)-1]] // 左边柱子
			//                                   // height: 右边柱子

			minHeight := min(left, height)
			width := i - stack[len(stack)-1] - 1
			area := (minHeight - middle) * width
			result += area

		}

		stack = append(stack, i)

	}

	return result

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// largestRectangleArea 柱状图中最大的矩形
// 找左右两边第一个比stack[len(stack)-1]矮的柱子
func largestRectangleArea(heights []int) int {

	// 收尾加0 ！！！
	// heights=[2,4,6,8]这种情况时，无法走到计算结果的逻辑！！
	// heights=[8,6,4,2]这种情况时，会导致无法凑齐3个元素
	heightsTemp := make([]int, 0, len(heights)+2)
	heightsTemp = append(heightsTemp, 0)
	heightsTemp = append(heightsTemp, heights...)
	heightsTemp = append(heightsTemp, 0)
	heights = heightsTemp
	fmt.Println(heights)

	stack := make([]int, 0, 10)

	maxArea := -1
	for i, height := range heights {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}

		for len(stack) > 0 && height < heights[stack[len(stack)-1]] {

			middle := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1] // 弹出middle中间这个柱子

			if len(stack) == 0 { // 找不到左边第一个比middle小的柱子
				break
			}

			width := i - stack[len(stack)-1] - 1 // 找到左边第一个比middle小的柱子
			area := middle * width
			if area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)

	}

	return maxArea

}
