package array

// minSubArrayLen 长度最小的子数组
func minSubArrayLen(target int, nums []int) int {

	minLen := len(nums) + 1
	sum := 0

	start := 0
	for end := 0; end < len(nums); end++ {
		sum += nums[end]

		for sum >= target {  // 这里是一个 while !!!!
			minLen = min(end - start + 1, minLen)

			// 移动滑动窗口的起始位置
			sum -= nums[start]     // start不会越过end,当start==end,sum=0,sum<正整数target
			start++
		}

	}



	if minLen == len(nums) + 1 {
		return 0
	} else {
		return minLen
	}

}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
} 