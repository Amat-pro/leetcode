package hash

// canConstruct 赎金信
func canConstruct(ransomNote string, magazine string) bool {
	record := [26]int{}
	for _, v := range magazine {
		record[v - rune('a')] ++
	}

	for _, v := range ransomNote {
		record[v - rune('a')] --
		if record[v - rune('a')] < 0 {
			return false
		}
	}

	return true
}