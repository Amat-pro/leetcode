package dp

import (
	"fmt"
	"testing"
)

func Test_bag01(t *testing.T) {

	weights := []int{1, 3, 4}
	values := []int{15, 20, 30}

	k := 4
	fmt.Println("dp bag01 result is: ", bag01(weights, values, k))

	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	fmt.Println("dp bag01 result is: ", bag_01(weight, value, 4))

	fmt.Println()
	fmt.Println("dp bag result is: ", bag(weight, value, 4))

}

// 01背包：  dp[i][j] = max(dp[i−1][j],dp[i−1][j−weights[i−1]]+values[i−1])
// 完全背包： dp[i][j] = max(dp[i-1][j], dp[i][j-weights[i-1]] + values[i-1])

func Test_lc_1143_longestCommonSubsequence(t *testing.T) {
	text1 := "abcde"
	text2 := "ace"

	fmt.Println("最长公共子序列长度, result is: ", longestCommonSubsequence(text1, text2)) // 3
}

func Test_lc_find_target_sum_ways(t *testing.T) {

	result := findTargetSumWays([]int{1, 1, 1, 1, 1}, 3)
	fmt.Println("目标和，result is: ", result)

}

func Test_lc_change(t *testing.T) {
	fmt.Println("零钱兑换II，result is: ", change([]int{1, 2, 5}, 5)) // 4
}

func Test_lc_coins_change(t *testing.T) {
	fmt.Println("零钱兑换，result is: ", coinChange([]int{1, 2, 5}, 11)) // 3
}

func Test_lc_0377_combination_sum4(t *testing.T) {
	fmt.Println("组合总和4-求排列数, result is: ", combinationSum4([]int{1, 2, 3}, 4)) // 7
}
