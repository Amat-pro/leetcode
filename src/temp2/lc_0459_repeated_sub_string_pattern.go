package temp

func repeatedSubStringPatter(s string) bool {

	next := getNext(s)

	if next[len(s)-1] == 0 {
		return false
	}

	if len(s)%(len(s)-next[len(s)-1]) == 0 {
		return true
	}

	return false

}

func kmpNext(s string) []int {
	next := make([]int, len(s))

	j := 0                        // 后缀的末尾
	for i := 1; i < len(s); i++ { // 前缀的末尾

		if s[i] != s[j] {
			j++
		} else {
			for j > 0 && s[i] != s[j] {
				j--
			}
		}

		next[i] = j

	}

	return next

}
