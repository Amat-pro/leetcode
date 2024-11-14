package array

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println("search result is: 4, ", search(nums, target))
}

func Test_removeElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println("removeElement result is: 2 ", removeElement(nums, val))
	fmt.Println(nums[0], nums[1])
}

func Test_sortedSquares(t *testing.T) {
	nums := []int{-7, -3, 2, 3, 11}
	fmt.Println("sortedSquares result is: ", sortedSquares(nums))
}

func Test_minSubArrayLen(t *testing.T) {
	target := 7
	nums := []int{2,3,1,2,4,3}
	fmt.Println("minSunArrayLen resule is: 2 ", minSubArrayLen(target, nums))
}
