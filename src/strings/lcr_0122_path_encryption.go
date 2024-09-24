package strings

// pathEncryption 路径加密 - 替换空格
func pathEncryption(path string) string {

	spaceCount := 0
	for _, v := range path {
		if v == ' ' {
			spaceCount++
		}
	}
	if spaceCount == 0 {
		return path
	}

	pathBytes := []byte(path)
	// 扩容2*space个位置
	temp := make([]byte, 2*spaceCount)
	left := len(pathBytes) - 1
	pathBytes = append(pathBytes, temp...)
	right := len(pathBytes) - 1

	for left >=0  {
		for left >= 0 && pathBytes[left] != ' ' { // 注意判断left的边界
			pathBytes[right] = pathBytes[left]
			left--
			right--
		}

		if left < 0 {  // 这里也要注意判断left的边界
			break
		}

		// 遇到pathBytes[left] == ''的情况了
		pathBytes[right] = '0'
		right--
		pathBytes[right] = '2'
		right--
		pathBytes[right] = '%'
		right--
		left--

	}

	return string(pathBytes)
}