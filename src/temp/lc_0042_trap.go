package temp

// trap 接雨水
func trap(h []int) int {

	if len(h) < 3 {
		return 0
	}

	stack := []int{} // 存放下标 - 单调递增栈
	result := 0 // 收集结果

	for i := 0; i < len(h); i ++ {

		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}

		var height int
		var width int
		for len(stack) > 1 && h[i] > h[stack[len(stack)-1]] { 
			// top作为中间柱子
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width = i - stack[len(stack)-1]-1
			height = min(h[i], h[stack[len(stack)-1]]) - h[top]
			result += width * height
		}

		stack = append(stack, i)
	}


	return result

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}