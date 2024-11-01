package dp

// longestCommonSubsequence 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {

	// dp数组
	// dp[i][j]: [0,i-1] text1 和 [0,j-1] text2 的最长公共子序列为dp[i][j]  求dp[len(text1)][len(text2)]
	// if text1[i-1] == text2[j-1] => dp[i][j] = dp[i-1][j-1] + 1
	// else dp[i][j] = max(dp[i-1][j], dp[i][j-1])

	// 初始化
	dp := make([][]int, 0, len(text1)+1)
	for j := 0; j < len(text1)+1; j++ {
		dp = append(dp, make([]int, len(text2)+1))
	}
	// dp[0][j] => 0
	// dp[i][0] => 0

	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {

			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxValue(dp[i-1][j], dp[i][j-1])
			}

		}
	}

	return dp[len(text1)][len(text2)]

}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
