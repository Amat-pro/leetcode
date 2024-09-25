package stack_queue

import "sort"

// topKFrequent 前k个高频元素
func topFrequent(nums []int, k int) []int {

	frequentMap := map[int]int{}
	for _, v := range nums {
		frequentMap[v]++
	}

	distinctNums := []int{}
	for num, _ := range frequentMap {
		distinctNums = append(distinctNums, num)
	}

	sort.Slice(distinctNums, func(a, b int) bool {

		// 频次按照从大到小排序
		if frequentMap[distinctNums[a]] > frequentMap[distinctNums[b]] {
			return true
		}

		return false

	})

	return distinctNums[:k]

}
