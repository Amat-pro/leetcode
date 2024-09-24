package strings

// reverseStr 翻转字符串II
func reverseStr(s string, k int) string {

	sbytes := []byte(s)
	for i := 0; i < len(sbytes); i += 2 * k {
		if i + k <= len(sbytes) {
			reverseString(sbytes[i:i+k])
			continue
		}
		reverseString(sbytes[i:])
	}

	return string(sbytes)

}