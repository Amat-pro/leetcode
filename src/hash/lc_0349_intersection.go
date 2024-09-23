package hash

// intersection 两个数组的交集
func intersection(nums1 []int, nums2[]int) []int {

	set := map[int]struct{}{}
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}

	result := []int{}
	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			result = append(result, v)
			delete(set, v) // 防止result结果集中出现重复元素
		}
	}

	return result

}