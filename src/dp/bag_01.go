package dp

import "fmt"

// bag01
// weights:
func bag01(weights []int, values []int, k int) int {
	// dp[i][j]: [0,i]物品，任放容量为j的背包
	// 不放物品i: dp[i-1][j]
	// 放物品i: dp[i-1][j-weights[i]] + values[i]

	dp := make([][]int, 0, len(weights))
	for i := 0; i <= k; i++ {
		dp = append(dp, make([]int, k+1))
	}

	// 初始化
	for i := 0; i < len(weights); i++ {
		dp[i][0] = 0
	}
	for i := 1; i <= k; i++ {
		if weights[0] <= i {
			dp[0][i] = values[0]
		}
	}

	for i := 1; i < len(weights); i++ { // 物品
		for j := 0; j <= k; j++ { // 背包
			if j < weights[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weights[i]]+values[i])
			}
		}
	}

	maxValue := 0
	for i := 0; i <= k; i++ {
		if dp[i][k] > maxValue {
			maxValue = dp[i][k]
		}
	}

	return maxValue

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// bag_01 01背包
func bag_01(weight, value []int, bagWeight int) int {
	// dp[j]: 容量为j的背包最大价值
	// 不放物品i: dp[j]
	// 放物品i: dp[j-weights[i]] + values[i]

	dp := make([]int, bagWeight+1)
	fmt.Println("start: ", dp)
	for i := 0; i < len(weight); i++ { // 物品
		for j := bagWeight; j >= weight[i]; j-- { // 背包   从大到小
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			fmt.Println("i: ", i, dp)
		}
	}

	return dp[bagWeight]
}

// bag 完全背包
func bag(weight, value []int, bagWeight int) int {
	// dp[j]: 容量为j的背包最大价值
	// 不放物品i: dp[j]
	// 放物品i: dp[j-weights[i]] + values[i]

	dp := make([]int, bagWeight+1)
	fmt.Println("start: ", dp)
	for i := 0; i < len(weight); i++ { // 物品
		for j := weight[i]; j <= bagWeight; j++ { // 背包  从小到大
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			fmt.Println("i: ", i, dp)
		}
	}

	return dp[bagWeight]
}
