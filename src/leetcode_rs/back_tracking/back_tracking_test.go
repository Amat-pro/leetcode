package back_tracking

import (
	"fmt"
	"testing"
)

func Test_combine(t *testing.T) {
	n := 4
	k := 2

	fmt.Println("combine result is: ", combine(n, k)) // [[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]
}

func Test_combinationSum3(t *testing.T) {
	k := 3
	n := 7

	fmt.Println("combinationSumIII result is: ", combinationSum3(k, n)) // [[1,2,4]]
}

func Test_letterCombinations(t *testing.T) {
	digits := "23"
	fmt.Println("letterCombinations result is: ", letterCombinations(digits)) // [ad ae af bd be bf cd ce cf]
}

func Test_combinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println("combinationSum result is: ", combinationSum(candidates, target)) // [[2 2 3] [7]]
}

func Test_combinationSum2(t *testing.T) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8

	fmt.Println("combinationSum2 result is: ", combinationSum2(candidates, target)) // [[1 1 6] [1 2 5] [1 7] [2 6]]
}

func Test_partition(t *testing.T) {
	s := "aab"
	fmt.Println("partition result is: ", partition(s)) // [[a a b] [aa b]]
}

func Test_restoreIpAddresses(t *testing.T) {
	s := "25525511135"
	fmt.Println("restoreIpAddresses result is: ", restoreIpAddresses(s)) // [255.255.11.135 255.255.111.35]
}

func Test_subSets(t *testing.T) {
	nums := []int{1, 2, 3}
	fmt.Println("subSets result is: ", subSets(nums)) // [[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]
}

func Test_subsetsWithDup(t *testing.T) {
	nums := []int{1, 2, 2}
	fmt.Println("subsetsWithSup result is: ", subsetsWithDup(nums)) // [[] [1] [1 2] [1 2 2] [2] [2 2]]
}

func Test_findSubsequences(t *testing.T) {
	nums := []int{4, 6, 7, 7}
	fmt.Println("findSubsequences result is: ", findSubsequences(nums)) // [[4 6] [4 6 7] [4 6 7 7] [4 7] [4 7 7] [6 7] [6 7 7] [7 7]]
}

func Test_permute(t *testing.T) {
	nums := []int{1, 2, 3}
	fmt.Println("permute result is: ", permute(nums)) // [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
}

func Test_permuteII(t *testing.T) {
	// nums := []int{1,1,2}
	// fmt.Println("permuteII result is: ", permuteII(nums)) // [[1 1 2] [1 2 1] [2 1 1]]
	nums := []int{1, 2, 3}
	fmt.Println("permuteII result is: ", permuteII(nums)) // [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
}
