package strings 

// kmp - strStr 找出字符串中第一个匹配项的下标
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return -1
	}

	next := next(needle)
	j := 0 // 遍历needle
	for i := 0; i < len(haystack); i ++ {
		// 不相等的情况
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] // 会推到next[j-1]位置重新匹配
		}

		if haystack[i] == needle[j] {
			j++
		}

		if j == len(needle) {  // 表示j已经匹配到了末尾 
			return i - len(needle) + 1
		}
	}

	return -1
}