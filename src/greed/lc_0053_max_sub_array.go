package greed

func maxSubArray(nums []int) int {

	if len(nums) == 0 {
		panic("invalid param")
	}

	maxSum := nums[0]
	sum := 0

	for i := 0; i < len(nums); i++ {

		sum += nums[i]
		if sum > maxSum {
			maxSum = sum
		}

		if sum < 0 {
			sum = 0
		}

	}

	return maxSum

}
