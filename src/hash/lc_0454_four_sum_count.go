package hash

// fourSumCount 四数之和II
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4[]int) int {
	
	// key: nums + nums2两数之和，value: count(两数之和出现的次数)
	map1 := map[int]int{}
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			map1[v1+v2]++
		}
	}

	result := 0
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			if _, ok := map1[0 - (v3+v4)]; ok {
				result += map1[0 - (v3+v4)]
			}
		}
	}

	return result

}