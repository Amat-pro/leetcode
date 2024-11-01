package greed

// widdleMaxLength - 0376 - 摆动序列
func widdleMaxLength(nums []int) int {

	if len(nums) == 0 {
		panic("invalid param")
	}
	if len(nums) <= 1 {
		return len(nums)
	}

	// 1. 正常情况
	//   2
	// /   \
	// 1    1(i)

	// 2. 平坡情况1
	//  2 - 2 - 2
	// /         \
	// 1          1(i)

	// 3. 平坡情况2（平坡后面是延续平坡前的单调递增或递减情况）
	//          4
	//         /
	//        3(i)
	//       /
	//  2 - 2
	// /
	// 1

	// prediff
	// curdiff
	// 情况1：当 (prediff > 0 && curdiff) < 0 || (prediff < 0 && curdiff > 0) => i满足一个摆动
	// 情况2：只记录最右侧的2 当 (prediff = 0 && curdiff) < 0 || (prediff = 0 && curdiff > 0) => i满足一个摆动
	//       只记录最左侧的2 当 (prediff > 0 && curdiff) = 0 || (prediff < 0 && curdiff = 0) => i满足一个摆动
	// 情况2中出现情况3时，会出现问题 => 情况3中，prediff的变化是：prediff>0,prediff=0,prediff>0,最右侧的2被错误算进去了
	// 面对情况3，不需要每次都更新prediff，只有摆动序列出现时才更新prediff(让prediff只代表一个趋势)

	result := 1 // 默认右侧有一个峰值

	prediff := 0 // 认为nums[0]左侧有一个平坡
	curdiff := 0

	for i := 0; i < len(nums)-1; i++ {
		curdiff = nums[i+1] - nums[i]
		if (prediff >= 0 && curdiff < 0) || (prediff <= 0 && curdiff > 0) {
			result++
			prediff = curdiff
		}
	}

	return result

}
