package dp

import "fmt"

// TODO

// fib 509 - 斐波那契数
func fib(n int) int {
	// dp[i] = dp[i-1] + dp[i-2]
	// 初始化：dp[0], dp[1]
	// 求dp[i]

	// 初始化
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	// 递推
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]

}

// climbStairs 70 - 爬楼梯
func climbStairs(n int) int {

	// dp[i]: 到i阶有多少种方法
	// 求dp[n]

	// 爬到i阶：
	// 从i-1阶爬上来
	// 从i-2阶爬上来
	// 递推公式： dp[i] = dp[i-1] + dp[i-2]

	// dp[0] = 0
	// dp[1] = 1
	// dp[2] = 2

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]

}

// minCostClimbingStairs 746 - 使用最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {

	// dp[i]: 表示爬到i阶需要的最小花费为dp[i]
	// 递推公式:
	// 爬到i阶：
	// 从dp[i-1]爬上来 => 花费cost[i-1]
	// 从dp[i-2]爬上来 => 花费cost[i-2]
	// dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	// 求dp[len(cost)]

	// dp[0] = 0
	// dp[1] = cost[0]

	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0 // 选择从1出发消耗最小为1

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	fmt.Println(dp)

	return dp[len(cost)]

}

// uniquePaths 62 - 不同路径
func uniquePaths(m int, n int) int {

	// dp[i][j]: 到达dp[i][j]有dp[i][j]种路径
	// 求dp[i][j]

	// dp[i][j]
	// 从dp[i][j-1]到达dp[i][j]
	// 从dp[i-1][j]到达dp[i][j]
	// => di[i][j] = dp[i][j-1] + dp[i-1][j]

	// 初始化 !!!
	// dp[0][0] = 0
	// dp[0][j] = 1  !!!
	// dp[i][0] = 1  !!!

	dp := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		temp := make([]int, n)
		dp = append(dp, temp)
	}
	// 这里的初始化，到dp[0][j]都是只有一条路径
	//             到dp[i][0]也都是只有一条路径（因为移动只能向下或者向右）
	for x := 0; x < n; x++ {
		dp[0][x] = 1
	}
	for x := 0; x < m; x++ {
		dp[x][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	return dp[m-1][n-1]

}

// uniquePathsWithObstacles - 63 - 不同路径II
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// obstacleGrid[i][j] == 1 => 是障碍物
	// 如果obstacleGrid[i][j] == 1 =》 dp[i][j] == 0 // 遇到障碍物，则从这里到下一地点的路径数量为0
	//    obstacleGrid[i][j] == 0 =》 dp[i][j] == dp[i][j] = dp[i][j-1] + dp[i-1][j]

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		temp := make([]int, n)
		dp = append(dp, temp)
	}

	for x := 0; x < n; x++ {
		dp[0][x] = 1
	}
	for x := 0; x < m; x++ {
		dp[x][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}

	return dp[m-1][n-1]

}

// integerBreak 343 - 整数拆分
func integerBreak(n int) int {
	// dp[i]: 拆分i，可以得到最大乘积dp[i]
	// i拆成2个: j * i-j (j: [1,i-1])
	// i拆成3+个: j * dp[i-j] (j: [1,i-1])
	// dp[i] = max(j*(i-j), j*dp[i-j], dp[i])
	// 求dp[n]

	dp := make([]int, n+1)
	dp[0] = 0 // 无意义
	dp[1] = 0 // 无意义
	dp[2] = 1 // (1*1)

	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(max(j*(i-j), j*dp[i-j]), dp[i])
		}
	}

	return dp[n]
}

// 0-1背包 - 物品i只有1个
// weights: 重量
// values: 价值
// k: 容量
func bag01(weights []int, values []int, k int) int {

	// dp[i][j]: [0,i]物品任选,容量为j的背包，最大价值为dp[i][j]
	// 对于物品i:
	// 1. 放物品i   => dp[i-1][j-weights[i]] + values[i] (这里放物品i,必须是容量为j-weights[i],[0,i-1]物品任选，因为物品i只有1个)
	// 2. 不放物品i => dp[i-1][j]
	// => dp[i][j] = max(dp[i-1][j-weights[i]]+values[i], dp[i-1][j])
	// 初始化:
	// dp[i][0] = 0
	// dp[0][j] = if weights[0] < j, dp[0][j]=values[0]
	// 求dp[len(weights)-1][k]

	dp := make([][]int, 0, len(weights))
	for x := 0; x < len(weights); x++ {
		temp := make([]int, k+1)
		dp = append(dp, temp)
	}
	// 初始化
	for j := 0; j <= k; j++ {
		if weights[0] <= j {
			dp[0][j] = values[0]
		}
	}

	for i := 1; i < len(weights); i++ { // 物品
		for j := 0; j <= k; j++ { // 背包
			if j < weights[i] {
				// 只能选择不放物品i
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j-weights[i]]+values[i], dp[i-1][j])
			}
		}

		fmt.Println(dp[i])

	}

	return dp[len(weights)-1][k]

}

// bag01_II 01背包 - 滚动数组
func bag01_II(weights []int, values []int, k int) int {

	// dp[j]: 容量i的背包最大价值为dp[i]
	// 放物品i: dp[j-weights[i]]+ values[i]
	// 不放物品i: dp[j]

	dp := make([]int, k+1)
	for i := 0; i < len(weights); i++ { // 物品
		for j := k; j >= weights[i]; j-- { // 背包 （从大到小） ！！！
			dp[j] = max(dp[j], dp[j-weights[i]]+values[i])
		}
		fmt.Println(dp)
	}

	return dp[k]

}

// bag - 完全背包 (物品i有无限个)
// weights: 重量
// values: 价值
// k: 容量
func bag(weights []int, values []int, k int) int {
	// dp[i][j]: [0,i]物品任选,容量为j的背包，最大价值为dp[i][j]
	// 对于物品i:
	// 1. 放物品i   => dp[i][j-weights[i]] + values[i] (这里放物品i,可以是容量为j-weights[i],[0,i]物品任选)
	// 2. 不放物品i => dp[i-1][j]
	// => dp[i][j] = max(dp[i][j-weights[i]]+values[i], dp[i-1][j])
	// 初始化:
	// dp[i][0] = 0
	// dp[0][j] = max(0, dp[0][j-weights[0]]+values[0])  // 不放物品0: 0； 放物品0: dp[0][j-weights[0]]+values[0]
	// 求dp[len(weights)-1][k]

	dp := make([][]int, 0, len(weights))
	for x := 0; x < len(weights); x++ {
		temp := make([]int, k+1)
		dp = append(dp, temp)
	}
	// 初始化
	for j := 0; j <= k; j++ {
		if j-weights[0] >= 0 {
			dp[0][j] = max(0, dp[0][j-weights[0]]+values[0]) // 不放物品0: 0； 放物品0: dp[0][j-weights[0]]+values[0]
		}
	}
	fmt.Println(dp[0])

	for i := 1; i < len(weights); i++ { // 物品
		for j := 0; j <= k; j++ { // 背包
			if j < weights[i] {
				// 只能选择不放物品i
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i][j-weights[i]]+values[i], dp[i-1][j])
			}
		}

		fmt.Println(dp[i])

	}

	return dp[len(weights)-1][k]
}

// bag_II - 完全背包 - 滚动数组
func bag_II(weights []int, values []int, k int) int {
	// dp[j]: 容量i的背包最大价值为dp[i]
	// 放物品i: dp[j-weights[i]]+ values[i]
	// 不放物品i: dp[j]

	dp := make([]int, k+1)
	for i := 0; i < len(weights); i++ { // 物品
		for j := weights[i]; j <= k; j++ { // 背包 （从小到大） ！！！
			dp[j] = max(dp[j], dp[j-weights[i]]+values[i])
		}
		fmt.Println(dp)
	}

	return dp[k]
}

// 分割等和子集
// 最后一块石头的重量II
// 目标和
// 一和零
// 零钱兑换II
// 零钱兑换
// 完全平方数
// 单词拆分

// 12
// 最长递增子序列
// 最长连续递增子序列
// 最长重复子数组
// 最长公共子序列
// 不相交的线
// 最大子序和
// 判断子序列
// 不同的子序列
// 两个字符串的删除操作
// 编辑距离
// 回文子串
// 最长回文子序列

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
