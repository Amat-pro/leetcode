package dp

// change 零钱兑换II 求组合数
func change(coins []int, amount int) int {

	// 完全背包
	// 组合数
	// dp[j]: 装满容量为j的背包有dp[j]种方法  求dp[amount]
	// dp[j] += dp[j-coins[i]]
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 0; i < len(coins); i++ {  // 遍历物品
		for j := coins[i]; j <= amount; j++ { // 遍历背包
			dp[j] += dp[j-coins[i]]
		}
	}

	return dp[amount]

}

