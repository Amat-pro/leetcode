package temp

func strStr(haystack, neele string) int {

	next := getNext(neele)
	j := 0                               // j表示后缀的末尾
	for i := 0; i < len(haystack); i++ { // i表示前缀的末尾
		if haystack[i] == neele[j] {
			j++
			if j == len(neele) {
				return i - len(neele) + 1
			}
		}

		for j > 0 && haystack[i] != neele[j] {
			j = next[j-1]
		}
	}

	return -1

}
