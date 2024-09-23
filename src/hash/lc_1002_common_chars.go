package hash

// commonChars 查找共用字符
func commonChars(words []string) []string {
	l := len(words)

	if l == 0 {
		return []string{}
	}

	record := [][26]int{}

	for i := 0; i < l; i++ {
		row := [26]int{}
		for _, r := range words[i] {
			row[r - rune('a')]++
		}
		record = append(record, row)
	}

	result := []string{}
	// 每一列的值中，最小值不为0，说明在每个字符串中都出现过
	for i := 0; i < 26; i++ {
		minVal := record[0][i]
		for j := 0; j < l; j++ {
			minVal = min(minVal, record[j][i])
		}
		if minVal == 0 {
			continue
		}

		char := string(rune(i) + 'a')
		// minVal表示共用的最小次数
		for i := 0; i < minVal; i++ {
			result = append(result, char)
		}
	}

	return result

}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}