package hash

// isAnagram 有效的字母异位词
func isAnagram(s string, t string) bool {

	record := [26]int{}
	for _, r := range s {
		record[r - rune('a')]++
	}
	for _, r := range t {
		record[r - rune('a')]--
	}

	return record == [26]int{}

	// for _, v := range record {
	// 	if v != 0 {
	// 		return false
	// 	}
	// }

	// return true
}