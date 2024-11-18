package dp

import (
	"fmt"
	"testing"
)

func Test_fib(t *testing.T) {
	n := 4
	fmt.Println("fib result is: ", fib(n)) // 3
}

func Test_climbStairs(t *testing.T) {
	n := 3
	fmt.Println("climbStairs result is: ", climbStairs(n)) // 3
}

func Test_minCostClimbingStairs(t *testing.T) {
	cost := []int{10, 15, 20}
	fmt.Println("minCostClimbingStairs result is: ", minCostClimbingStairs(cost)) // 15
	cost2 := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println("minCostClimbingStairs result is: ", minCostClimbingStairs(cost2)) // 6
}

func Test_uniquePaths(t *testing.T) {
	m := 3
	n := 7
	fmt.Println("uniquePaths result is: ", uniquePaths(m, n))
}

func Test_uniquePathsWithObstacles(t *testing.T) {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	fmt.Println("uniquePathsWithObstacles result is: ", uniquePathsWithObstacles(obstacleGrid)) // 2
}

func Test_integerBreak(t *testing.T) {
	n := 10
	fmt.Println("整数拆分 - integerBreak result is: ", integerBreak(n)) // 36
}

func Test_bag01(t *testing.T) {
	weights := []int{1, 3, 4}
	values := []int{15, 20, 30}

	k := 4

	fmt.Println("bag01 result is: ", bag01(weights, values, k))
	fmt.Println("滚动数组-bag01II result is: ", bag01_II(weights, values, k))
}

func Test_bag(t *testing.T) {
	weights := []int{1, 3, 4}
	values := []int{15, 20, 30}

	k := 4

	fmt.Println("二维数组-bagII result is: ", bag(weights, values, k))
	fmt.Println("滚动数组-bag_II result is: ", bag_II(weights, values, k))
}

func Test_canPartition(t *testing.T) {
	nums := []int{1, 5, 11, 5}
	fmt.Println("canPartition result is: ", canPartition(nums)) // true
}

func Test_lastStoneWeightII(t *testing.T) {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println("lastStoneWeightII result is: ", lastStoneWeightII(stones)) // 1
}

func Test_findTargetSumWays(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	fmt.Println("findTargetSumWays result is: ", findTargetSumWays(nums, target)) // 5
}

func Test_change(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 5
	fmt.Println("零钱兑换II-change result is: ", change(amount, coins)) // 4
}

func Test_coinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println("零钱兑换-coinChange result is: ", coinChange(coins, amount)) // 3
}

func Test_numSquares(t *testing.T) {
	fmt.Println("numSquares result is: ", numSquares(13)) // 2
}

func Test_wordBreak(t *testing.T) {
	fmt.Println("单词拆分 - wordBreak result is: ", wordBreak("applepenapple", []string{"apple", "pen"}))
}

func Test_lengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println("最长递增子序列：lengthOfLIS result is: ", lengthOfLIS(nums)) // 4
}

func Test_findLengthOfLCIS(t *testing.T) {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println("最长连续递增子序列：lengthOfLCIS result is: ", findLengthOfLCIS(nums)) // 3
}

func Test_maxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("maxSubArray result is: ", maxSubArray(nums))       // 6
	fmt.Println("maxSubArray_II result is: ", maxSubArray_II(nums)) // 6

	nums1 := []int{5, 4, -1, 7, 8}
	fmt.Println("maxSubArray result is: ", maxSubArray(nums1))       // 23
	fmt.Println("maxSubArray_II result is: ", maxSubArray_II(nums1)) // 23
}
