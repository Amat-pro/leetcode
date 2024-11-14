package hash

import (
	"fmt"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	s := "anagram"
	s2 := "nagaram"
	fmt.Println("isAnagram result is: true ", isAnagram(s, s2))
}

func Test_commonChars(t *testing.T) {
	words := []string{"bella", "label", "roller"}
	fmt.Println("commonChars result is: ell ", commonChars(words))
}

func Test_intersection(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Println("intersection result is: 2 ", intersection(nums1, nums2))

}

func Test_isHappy(t *testing.T) {
	fmt.Println("isHappy result is: true ", isHappy(19))
}

func Test_twoSum(t *testing.T) {
	fmt.Println("twoSum result is: [0,1] ", twoSum([]int{2, 7, 11, 15}, 9))
}

func Test_fourSumCount(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}

	fmt.Println("fourSumCount result is: 2 ", fourSumCount(nums1, nums2, nums3, nums4))
}

func Test_canConstruct(t *testing.T) {
	fmt.Println("canConstruct result is: true ", canConstruct("aa", "aab"))
}
