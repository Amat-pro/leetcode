package strings

// reverseString 反转字符串
func reverseString(s []byte) {

	left := 0
	right := len(s) - 1
	for ; left < len(s) / 2; {
		s[left],s[right] = s[right],s[left]
		left++
		right--

	} 

}