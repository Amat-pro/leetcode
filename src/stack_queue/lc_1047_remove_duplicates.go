package stack_queue

// removeDuplicates 删除字符串中的所有相邻重复项
func removeDuplicates(s string) string {

	var stack []byte
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && s[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)

}
