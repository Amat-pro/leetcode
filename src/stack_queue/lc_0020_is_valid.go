package stack_queue

// isValid 有效的括号
func isValid(s string) bool {

	if s == "" {
		return false
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '[' || s[i] == '(' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if (s[i] == ']' && s[len(stack)-1] == '[') ||
			(s[i] == ')' && s[len(stack)-1] == '(') ||
			(s[i] == '}' && s[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return true

}
