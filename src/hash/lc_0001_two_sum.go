package hash

// twoSum 两数之和
func twoSum(nums []int, target int) []int {
	set := map[int]int{}

	for i, v := range nums {
		if _, ok := set[target - v]; ok {
			return []int{i, set[target -v]}
		} else {
			set[v] = i
		}
	}

	return []int{}
}