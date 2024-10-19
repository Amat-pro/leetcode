package temp

// dailyTemperatures 每日温度
func dailyTemperatures(temperatures []int) []int {

	if len(temperatures) == 0 {
		return []int{}
	}

	result := make([]int, len(temperatures)) // 收集结果,初始化为0
	stack := []int{}                         // 记录下标

	stack = append(stack, 0)
	for i := 1; i < len(temperatures); i++ {
		// 当前遍历元素比较栈顶元素
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1] // 出栈
		}

		// 当前元素入栈
		stack = append(stack, i)
	}

	return result

}
