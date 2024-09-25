package strings

// repeatedSubStringPattern 重复的子字符串
func repeatedSubStringPattern(s string) bool {
	l := len(s)

	if l == 0 {
		return false
	}

	next := next(s)
	if next[l-1] == 0 {
		return false
	}

	if l % (l - next[l-1]) == 0 {
		return true
	}

	return false
}