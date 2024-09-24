package two_point

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println("three sum, result is: ", threeSum(nums))
}

func TestFourSum(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println("four sum, result is: ", fourSum(nums, target))
}
