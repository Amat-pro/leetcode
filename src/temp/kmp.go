package temp

// func getNext(s []byte) []int {

// 	// next数组表示：next[i]表示i结尾的字符串它的最长相等前后缀的长度

// 	next := make([]int, len(s))

// 	j := 0 // 前缀末尾位置
// 	for i := 1; i < len(s); i++ { // 表示后缀末尾位置
// 		if s[i] == s[j] {
// 			j++
// 		} else {
// 			for j > 0 && s[i] != s[j] {
// 				j--
// 			}
// 		}

// 		next[i] = j
// 	}

// 	return next


// }

func getNext(b []byte) []int {

	// next数组：next[i]结尾的最长相等前后缀长度
	next := make([]int, len(b))

	j := 0 // j表示后缀的末尾
	for i := 1; i < len(b); i++ { // i表示前缀的末尾
		if b[i] == b[j] {
			j++
		} else {
			for j > 0 && b[i] != b[j] {
				j--
			}
		}

		next[i] = j
	}

	return next

}