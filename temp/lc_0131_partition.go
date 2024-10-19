package temp

func partition(s string) [][]string {

	result := [][]string{}

	path := []string{}

	var backTracking func(index int)
	backTracking = func(index int) {

		// 确定终止条件
		if index == len(s) {
			for _, v := range path {
				if !isHuiwen([]byte(v)) {
					return
				}
			}
			// 收集结果
			temp := make([]string, len(path))
			copy(temp, path)
			result = append(result, temp) 
			// 写成 result = append(result, path) 是不对的 ！！！  【path会在回溯中被修改！！！】
		}

		// 单层处理逻辑
		for i := index; i < len(s); i++ {
			tempStr := s[index : i+1]
			path = append(path, tempStr)
			// 递归
			backTracking(i + 1)
			// 回溯
			path = path[:len(path)-1]
		}

	}

	backTracking(0)
	return result

}

func isHuiwen(b []byte) bool {

	if len(b) == 1 {
		return true
	}

	left := 0
	right := len(b) - 1

	for left < right {
		if b[left] != b[right] {
			return false
		}

		left++
		right--
	}

	return true
}