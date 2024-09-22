package array

import (
	"fmt"
	"testing"
)

func Test_lc_0027_remove_element(t *testing.T) {
	nums := []int{1,2,3,4,5}
	val := 3

	k := removeElements(nums, val)
	fmt.Println("remove element, result is: ", k, nums)

}

func TestSortedSquares(t *testing.T) {
	nums := []int{-4, -1, 0, 3, 10}
	result := sortedSquares(nums)
	fmt.Println("sorted squares, result is: ", result)
}

func TestMinSubArrayLen(t *testing.T) {
	nums := []int{2,3,1,2,4,3}
	target := 7
	fmt.Println("min sub array, result is: ", minSubArrayLen(target, nums)) // 2


	nums1 := []int{100}
	target1 := 101
	fmt.Println("min sub array, result is: ", minSubArrayLen(target1, nums1)) // 0
}

func TestGenerateMatrix(t *testing.T) {
	fmt.Println("generate matrix, n=3, result is: ", generateMatrix(3))
	fmt.Println("generate matrix, n=4, result is: ", generateMatrix(4))
}