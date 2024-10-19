package strings

// next 获取next数组
func next(s string) []int {

	byteS := []byte(s)
	l := len(byteS)
	if l == 0 {
		return []int{}
	}

	next := make([]int, l)
	next[0] = 0

	j := 0 //j指向前缀的末尾
	for i := 1; i < l; i++ { // i指向后缀的末尾
		for j > 0 && byteS[i] != byteS[j] {
			j = next[j-1]
		}

		if byteS[i] == byteS[j] {
			j += 1
		}

		next[i] = j
	}

	return next

}