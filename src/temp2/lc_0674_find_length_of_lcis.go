package temp

// findLengthOfLCIS 最长连续递增序列
func findLengthOfLCIS(nums []int) int {

	// dp[i]: 以i结尾的数组，最长连续递增序列为dp[i]
	// dp[i]:
	// nums[i] > nums[i-1] dp[i]=dp[i-1]+1
	// nums[i] <= nums[i-1] dp[i]=1

	// i: [0, len(nums)-1]
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	maxLength := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
			if dp[i] > maxLength {
				maxLength = dp[i]
			}
		} else {
			dp[i] = 1
		}
	}

	return maxLength

}
