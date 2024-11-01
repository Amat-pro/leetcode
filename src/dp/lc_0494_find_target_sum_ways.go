package dp

// findTargetSumWays 目标和
func findTargetSumWays(nums []int, target int) int {

	// 0/1背包 =》倒叙遍历背包容量

	// left(正数集合) + right(负数集合) = sum
	// left - right = target
	// left = (target+sum) / 2   【left如果不为整数，说明没有这样的组合】
	// 求装满容量为left的背包有多少种方法

	// dp[j]: 装满容量为j的背包有dp[j]种方法
	// dp[j] += dp[j-coins[i]]

	// 例子：
	// dp[5]: 装满dp[5]的背包有dp[5]种方法
	// 已有
	// 1    dp[4]种方法可以凑成dp[5]  => 已有大小为1的物品时，再有dp[4]种方法就可以凑成dp[5]
	// 2    dp[3]种方法可以凑成dp[5]
	// 3    dp[2]
	// 4    dp[1]
	// 5    dp[0]
	// 所以凑成dp[5]有累加起来种方法

	// 初始化
	// dp[0] = 1

	// 遍历顺序
	// 组合数 or 排列数 ？ =》 组合数，不需要区分数序

	sum := 0
	for _, v := range nums {
		sum += v
	}
	bagSize := (sum + target) / 2

	dp := make([]int, bagSize+1)
	dp[0] = 1

	// 遍历背包
	for i := 0; i < len(nums); i++ { // 遍历物品
		for j := bagSize; j >= nums[i]; j-- { // 遍历背包 - 倒序
			dp[j] += dp[j-nums[i]]
		}

	}

	return dp[bagSize]

}
