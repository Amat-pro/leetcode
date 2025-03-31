# leetcode

## TODO
回溯
DP

## 其他记录
[一刷](https://github.com/Amat-pro/cargo_demo/blob/develop/leetcode/index.md)

## 再刷
### 数组
[dir - array](./src/leetcode_rs/array/array.go)

### cache
[dir - cache](./src/leetcode_rs/cache/lru_cache.go)

### 哈希表
[dir - hash](./src/leetcode_rs/hash/hash.go)

### 字符串
[dir - string](./src/leetcode_rs/string/string.go)

### 双指针
[dir - two_point](./src/leetcode_rs/two_point/twp_point.go)

### 排序
[dir - sort](./src/leetcode_rs/sort/mod.rs)

### 栈/队列/单调栈
[dir - queue/stack](./src/leetcode_rs/queue_stack/queue_stack.go)

### nsp
[dir - nsp](./src/nsp/nsp.go)

### Tree - 二叉树
[dir - tree](./src/leetcode_rs/tree/tree.go)

### 回溯
[dir - back_tracking](./src/leetcode_rs/back_tracking/back_tracking.go)

### 动态规划
[dir - DP](./src/leetcode_rs/dp/dp.go)

### DP - 背包问题
```

通过二维数组的推导，可以看出01背包在使用一维数组求解时遍历背包是从大到小的顺序，完全背包在使用一维数组求解时遍历背包是从小到大的顺序。本质上是递推公式差异。

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
```

//