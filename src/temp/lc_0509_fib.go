package temp

func fib(n int) int {

	// dp[i]: 第i个值为i
	// 递推公式: dp[i] = dp[i-1] + dp[i-2]
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]

}