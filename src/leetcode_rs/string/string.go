package string

// reverseString 344 - 反转字符串
func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}

	left := 0
	right := len(s) - 1

	for left < right {
		s[left], s[right] = s[right], s[left]

		left++
		right--
	}
}

// reverseStr 541 - 翻转字符串II
func reverseStr(s string, k int) string {
	b := []byte(s)
	for i := 0; i < len(b); i += 2 * k {
		if i+k < len(b) {
			reverseString(b[i : i+k])
		} else {
			reverseString(b[i:])
		}
	}

	return string(b)
}

// reverseWords 反转字符串中的单词
func reverseWords(s string) string {
	b := removeSpace([]byte(s))

	// 反转整个字符串
	reverseString(b)
	// 反转每个单词
	start := 0
	for end := 0; end <= len(b); end++ {
		if end == len(b) || b[end] == ' ' { // end到末尾了或者end指向空格，则[start,end)是需要反转的单词
			reverseString(b[start:end])
			start = end + 1
		}
	}

	return string(b)
}

func removeSpace(b []byte) []byte {
	// 删除前边空格
	// 删除中间多余空格
	// 删除尾部空格

	slow := 0
	fast := 0
	for fast < len(b) && b[fast] == ' ' { // 移除前边的空格
		fast++
	}

	for fast < len(b) {
		if fast-1 > 0 && b[fast] == b[fast-1] && b[fast] == ' ' { // 保留每个单词后面第一个空格
			fast++
			continue
		}
		b[slow] = b[fast]
		slow++
		fast++
	}

	// 移除尾部的空格
	// slow-1是末尾位置（slow指向的是下一个写的位置，可能是等于length）
	if b[slow-1] == ' ' {
		return b[:slow-1]
	}
	return b[:slow]

}

// strStr 28 - 找出字符串中第一个匹配项的下标
func strStr(haystack string, needle string) int {
	next := getNext([]byte(needle))

	j := 0                               // j指向后缀的末尾
	for i := 0; i < len(haystack); i++ { // i指向前缀的末尾

		// 相等的情况
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) { // 已经匹配到了末尾
			return i - len(needle) + 1
		}

		// 不想等的情况
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}

	}

	return 0
}

func getNext(b []byte) []int {
	next := make([]int, len(b))
	next[0] = 0

	j := 0                        // j指向前缀的末尾
	for i := 1; i < len(b); i++ { // i从1开始 i指向后缀的末尾
		if b[i] == b[j] { // 相等的情况
			j++
		} else {
			for j > 0 && b[i] != b[j] { // 不想等的情况
				j = next[j-1]
			}
		}
		next[i] = j
	}

	return next
}

// repeatedSubstringPattern 459 - 重复的子字符串 - KMP解法
func repeatedSubstringPattern(s string) bool {
	// 最长公共字符串不包含的部分 -> 重复单元
	next := getNext([]byte(s))

	if next[len(s)-1] == 0 {
		return false
	}

	if len(s)%(len(s)-next[len(s)-1]) == 0 {
		return true
	} else {
		return false
	}

}
