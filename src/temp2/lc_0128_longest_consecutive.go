package temp

// longestConsecutive 最长连续序列
func longestConsecutive(nums []int) int {

	// dp[i]: [0, i]范围内nums最长连续序列为dp[i]
	// [0, i-1]范围内(temp), nums[i] > nums[temp] dp[i]=max(dp[i],dp[i]+1)

	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	maxLength := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				if dp[i] > maxLength {
					maxLength = dp[i]
				}
			}

		}

	}

	return maxLength

}
