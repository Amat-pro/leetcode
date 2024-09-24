package strings

// reverseWords 翻转单词
func reverseWords(s string) string {

	b := []byte(s)
	if len(b) == 0 {
		return s
	}

	// 移除前面，中间，末尾多余的空格
	slow := 0
	for fast := 0; fast < len(b); fast++ {

		if b[fast] == ' ' {
			continue
		}

		if slow != 0 { // 每个单词开头是空格
			b[slow] = ' '
			slow++
		}

		for fast < len(b) && b[fast] != ' ' {
			b[slow] = b[fast]
			fast++
			slow++ // slow做了++, 是末尾+1的位置
		}

	}

	b = b[:slow] // slow是末尾节点+1的位置

	// 翻转整个字符串
	reverseString(b)

	// 翻转每个单词
	start := 0
	for i := 0; i < len(b); i++ {
		// 遇到了一个单词结尾处的空格
		if b[i] == ' ' { 
			reverseString(b[start:i])
			start = i + 1
			continue
		} 
		
		if i == len(b) - 1 {
			// i指向了末尾
			reverseString(b[start:i+1])
			break
		}
	}

	return string(b)

}
