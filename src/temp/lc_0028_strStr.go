package temp

func strStr(haystack string, needle string) int {

	next := getNext([]byte(needle))

	j := 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[j] {
			j++

			if j == len(needle) {
				return i - len(needle) + 1
			}

		}

		for j > 0 && needle[j] != haystack[i] {
			j = next[j-1]
		}
	}

	return -1

}