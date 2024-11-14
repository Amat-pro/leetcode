package hash

import (
	"math"
)

// isAnagram 有效的字母异位词
func isAnagram(s string, t string) bool {

	// 26个字母
	record := make([]int, 26)
	for _, c := range s {
		record[c-rune('a')]++
	}
	for _, c := range t {
		record[c-rune('a')]--
	}

	for _, v := range record {
		if v != 0 {
			return false
		}
	}

	return true

}

// commonChars 查找公用字符
func commonChars(words []string) []string {
	// 二维数组记录每个字符串中字符的频率
	records := make([][]int, 0, len(words))
	for i := 0; i < len(words); i++ {
		record := make([]int, 26)
		records = append(records, record)
	}

	for i, word := range words {
		for _, c := range word {
			records[i][c-'a']++
		}
	}

	result := []string{}
	for i := 0; i < 26; i++ {
		minValue := math.MaxInt // 每个数组中找到频率最低的
		for _, record := range records {
			minValue = min(minValue, record[i])
		}
		if minValue > 0 {
			for count := 0; count < minValue; count++ {
				result = append(result, string(rune(i)+'a'))
			}
		}
	}

	return result

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// intersection 349 -两个数组的交集
func intersection(nums1 []int, nums2 []int) []int {
	hash := map[int]int{}
	for _, v := range nums1 {
		if _, ok := hash[v]; !ok {
			hash[v] = v
		}
	}

	result := []int{}
	for _, v := range nums2 {
		num1, ok := hash[v]
		if ok {
			result = append(result, num1)
			// 去重
			delete(hash, num1)
		}
	}

	return result
}

// isHappy 202 - 快乐数
func isHappy(n int) bool {

	// 题目中说可能会无限循环，说明每一次的n值会重复，当n重复时则说明进入了无限循环，不会有结果
	// 这里
	nHash := map[int]bool{}

	for !nHash[n] {
		sum := getNum(n)
		nHash[n] = true // n已经出现过

		n = sum
	}

	return n == 1

}

func getNum(n int) int {
	sum := 0
	for n != 0 {
		temp := n % 10
		sum += temp * temp

		n = n / 10

	}

	return sum
}

// twoSum 两数之和
func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for i, v := range nums {
		if index, ok := hash[target-v]; ok {
			return []int{index, i}
		}
		hash[v] = i
	}

	return nil

}

// fourSumCount 454 - 四数相加
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	hash := map[int]int{} // 记录两数之和的频率
	for _, i := range nums1 {
		for _, j := range nums2 {
			hash[i+j]++
		}
	}

	result := 0
	for _, i := range nums3 {
		for _, j := range nums4 {
			count, ok := hash[-i-j]
			if ok {
				result += count
			}
		}
	}

	return result

}

// canConstruct 赎金信
func canConstruct(randomNote string, magazine string) bool {
	record := make([]int, 26)
	for _, v := range magazine {
		record[v-'a']++
	}

	for _, v := range randomNote {
		record[v-'a']--
		if record[v-'a'] < 0 {
			return false
		}
	}

	return true

}