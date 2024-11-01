package greed

import (
	"fmt"
	"testing"
)

func Test_lc_0455_find_content_children(t *testing.T) {

	g := []int{1, 2} // 饼干
	children := []int{1, 2, 3}
	fmt.Println("贪心-分发饼干，result is: ", findContentChildren(g, children)) // 2

}

func Test_lc_0376_wiggle_max_length(t *testing.T) {
	nums := []int{1, 7, 4, 9, 2, 5}
	fmt.Println("贪心-摆动序列，result is: ", widdleMaxLength(nums)) // 6
}

func Test_lc_0053_max_sub_array(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("贪心-最大子数组和，result is: ", maxSubArray(nums)) // 6
}
