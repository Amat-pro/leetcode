package strings


// reverseLeftWords 左旋转字符串
func reverseLeftWords(s string, n int) string {

	b := []byte(s)
	// 翻转[0,n)
	reverseString(b[0:n])
	// 翻转[n:l)
	reverseString(b[n:])
	// 翻转整个字符串
	reverseString(b)
	
	return string(b)
}