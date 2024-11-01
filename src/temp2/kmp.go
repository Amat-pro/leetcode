package temp

func getNext(s string) []int {

	next := make([]int, len(s))

	j := 0                        // j表示前缀的末尾
	for i := 1; i < len(s); i++ { // i表示后缀的末尾

		if s[i] == s[j] {
			j++
			next[i] = j
		} else {
			for j > 0 && s[i] != s[j] {
				j--
			}

			next[i] = j
		}

	}

	return next
}
