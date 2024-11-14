package two_point

import (
	"fmt"
	"testing"
)

func Test_ThreeSum(t *testing.T) {

	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println("threeSum result is: ", threeSum(nums)) // [[-1,-1,2],[-1,0,1]]

}

func Test_FourSum(t *testing.T) {

	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println("fourSum result is: ", fourSum(nums, target)) // [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

}
