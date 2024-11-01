package dp

import "math"

// coinChange 零钱兑换 - 完全背包问题 - 求最少物品数
func coinChange(coins []int, amount int) int {
	// dp[j]: 凑成j所需要的最少物品为dp[j]  求dp[amount]
	// dp[j] = min(dp[j], dp[j-coins[i]]+1)

	dp := make([]int, amount+1)
	dp[0] = 0
	for j := 1; j < amount+1; j++ {
		dp[j] = math.MaxInt
	}

	for i := 0; i < len(coins); i++ { // 遍历物品
		for j := coins[i]; j <= amount; j++ { // 遍历背包 - 正序
			if dp[j-coins[i]] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt {
		return -1
	}

	return dp[amount]

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
