package stack_queue

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	s := "([])"
	fmt.Println("is valid, result is: ", isValid(s)) // true
}

func TestRemoveDuplicates(t *testing.T) {
	s := "abbaca"
	fmt.Println("remove duplicates, result is: ", removeDuplicates(s))
}

func TestEvalRPN(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println("evalRPN, result is: ", evalRPN(tokens)) // 9
}

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println("max sliding window, result is: ", maxSlidingWindow(nums, 3)) // 3 3 5 5 6 7
}

func TestTopKFrequent(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println("top k frequent, result is: ", topFrequent(nums, k)) // 1 2
}
