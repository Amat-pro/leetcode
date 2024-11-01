package dp

func combinationSum4(nums []int, target int) int {

	// 求排列数 - 强调元素顺序
	// dp[j]: 装满容量为j的背包有dp[j]种组合  求dp[target]
	// dp[j] += dp[j-coins[i]]

	dp := make([]int, target+1)
	dp[0] = 1

	for j := 0; j <= target; j++ { // 先遍历背包
		for i := 0; i < len(nums); i++ {

			if nums[i] <= j {
				dp[j] += dp[j-nums[i]]
			}
		}

	}

	return dp[target]

}

// 求组合数：
// 遍历物品
// 遍历背包

// 求排列数：
// 遍历背包
// 遍历物品
