package temp

func repeatedSubStringPattern(s string) bool {

	// next数组
	next := getNext([]byte(s))

	// 最长相等前后缀-不包含的部分-是重复单元
	repeatedSubStrLen := len(s) - next[len(s)-1]
	if len(s) % repeatedSubStrLen == 0 {
		return true
	}

	return false

}