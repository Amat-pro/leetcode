package queue_stack

import (
	"fmt"
	"testing"
)

func Test_isValid(t *testing.T) {
	s := "()]{}"
	fmt.Println("isValid result is: ", isValid(s)) // false
}

func Test_removeDuplicates(t *testing.T) {
	s := "abbaca"
	fmt.Println("removeDuplicates, result is: ", removeDuplicates(s)) // ca
}

func Test_evalRPN(t *testing.T) {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	// tokens := []string{"2","1","+","3","*"} // 9
	fmt.Println("evalRPN, result is: ", evalRPN(tokens)) // 22
}

func Test_maxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println("maxSlidingWindow, result is: ", maxSlidingWindow(nums, k)) // [3,3,5,5,6,7]
}

func Test_dailyTemperatures(t *testing.T) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println("dailyTemperatures, result is: ", dailyTemperatures(temperatures)) // [1,1,4,2,1,1,0,0]
}

func Test_nextGreaterElement(t *testing.T) {
	// nums1 := []int{4, 1, 2}
	// nums2 := []int{1, 3, 4, 2}
	// fmt.Println("nextGreaterElement result is: ", nextGreaterElement(nums1, nums2)) // [-1,3,-1]
	nums1 := []int{2, 4}
	nums2 := []int{1, 2, 3, 4}
	fmt.Println("nextGreaterElement result is: ", nextGreaterElement(nums1, nums2)) // [3,-1]
}

func Test_nextGreaterElementII(t *testing.T) {
	nums := []int{1, 2, 1}
	fmt.Println("nextGreaterElementII result is: ", nextGreaterElementsII(nums)) // [2,-1,2]
}

func Test_trap(t *testing.T) {
	heights := []int{4, 2, 0, 3, 2, 5}
	fmt.Println("trap result is: ", trap(heights)) // 9
	heights2 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println("trap result is: ", trap(heights2)) // 6
}

func Test_largestRectangleArea(t *testing.T) {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println("largestRectangleArea result is: ", largestRectangleArea(heights)) // 10
	heights2 := []int{2, 4}
	fmt.Println("largestRectangleArea result is: ", largestRectangleArea(heights2)) // 4
}
