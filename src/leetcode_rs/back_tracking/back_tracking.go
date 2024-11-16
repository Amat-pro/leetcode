package back_tracking

import (
	"sort"
	"strconv"
	"strings"
)

// combine 77 - 组合
func combine(n int, k int) [][]int {

	result := make([][]int, 0, 10)
	path := make([]int, 0, 10)

	var backTracking func(n int, k int, startIndex int)
	backTracking = func(n int, k int, startIndex int) {

		// 终止条件 - 收集结果
		if len(path) == k {
			pathTemp := make([]int, k)
			copy(pathTemp, path)
			result = append(result, pathTemp)
		}

		// 遍历
		for i := startIndex; i <= n; i++ {

			// 剩余元素个数 < 需要的元素个数
			if n-i+1 < k-len(path) { // 剪枝操作
				break
			}

			path = append(path, i)
			backTracking(n, k, i+1)   // 递归
			path = path[:len(path)-1] // 回溯
		}

	}

	backTracking(n, k, 1)

	return result

}

// combinationSum3 - 216 组合总和III
// 固定数组1--9, 不需要排序
func combinationSum3(k int, n int) [][]int {

	result := make([][]int, 0, 10)
	path := make([]int, 0, k)

	var backTracking func(k int, n int, startIndex int, sum int)
	backTracking = func(k int, n int, startIndex int, sum int) {

		// 终止条件
		if len(path) == k {
			if sum == n {
				pathTemp := make([]int, k)
				copy(pathTemp, path)
				result = append(result, pathTemp)
			}

			return
		}

		for i := startIndex; i <= 9; i++ {

			path = append(path, i)
			sum += i

			backTracking(k, n, i+1, sum) // 递归

			path = path[:len(path)-1] // 回溯
			sum -= i                  // 回溯

		}

	}

	backTracking(k, n, 1, 0)

	return result
}

// letterCombinations 17 - 电话号码的字母组合
func letterCombinations(digests string) []string {

	numMap := map[string][]rune{
		"2": []rune("abc"),
		"3": []rune("def"),
		"4": []rune("ghi"),
		"5": []rune("jkl"),
		"6": []rune("mno"),
		"7": []rune("pqrs"),
		"8": []rune("tuv"),
		"9": []rune("wxyz"),
	}

	result := make([]string, 0, 10)
	path := make([]rune, 0, 10)

	var backTracking func(digestsIndex int)
	backTracking = func(digestsIndex int) {

		// 终止条件
		if digestsIndex == len(digests) {
			temp := string(path)
			result = append(result, temp)
			return
		}

		// 处理每一层 - 循环
		num := string(digests[digestsIndex])
		characters := numMap[num]
		for i := 0; i < len(numMap[num]); i++ {
			path = append(path, characters[i])
			backTracking(digestsIndex + 1) // 递归
			path = path[:len(path)-1]      // 回溯
		}

	}

	backTracking(0)

	return result

}

// combinationSum 39 - 组合总和
// candidates: 无重复元素的整数数组 (无重复-不需要考虑去重)
// target: 目标整数（目标和）
func combinationSum(candidates []int, target int) [][]int {

	result := make([][]int, 0, 10)
	path := make([]int, 0, 10)

	var backTracking func(startIndex int, sum int)
	backTracking = func(startIndex int, sum int) {
		// 终止条件
		if sum > target {
			return
		}
		if sum == target {
			// 收集结果
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		// 单层处理逻辑
		for i := startIndex; i < len(candidates); i++ {
			path = append(path, candidates[i])
			sum += candidates[i]

			backTracking(i, sum) // 递归 (每个元素可以取重复取，所以下一层中也可以使用i这个元素)

			// 回溯
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}

	backTracking(0, 0)

	return result

}

// combinationSum2 40 - 组合总和II - 解集不能包含重复的组合 && candidates包含重复元素
func combinationSum2(candidates []int, target int) [][]int {

	// 排序 ！！！
	sort.Slice(candidates, func(x, y int) bool {
		return candidates[x] < candidates[y]
	})

	result := make([][]int, 0, 10)
	path := make([]int, 0, 10)

	var backTracking func(startIndex int, sum int, used map[int]bool)
	backTracking = func(startIndex int, sum int, used map[int]bool) {

		// 终止条件
		if sum > target {
			return
		}
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		// 单层处理逻辑
		for i := startIndex; i < len(candidates); i++ {
			if i > 0 && candidates[i] == candidates[i-1] && !used[i-1] { // uses[i-1]在这层的处理中已经被使用过
				continue
			}

			path = append(path, candidates[i])
			sum += candidates[i]
			used[i] = true

			backTracking(i+1, sum, used) // 递归

			// 回溯
			path = path[:len(path)-1]
			sum -= candidates[i]
			used[i] = false

		}

	}

	used := map[int]bool{}
	backTracking(0, 0, used)

	return result

}

// partition 131 - 分割回文串
func partition(s string) [][]string {
	result := make([][]string, 0, 10)
	path := make([]string, 0, 10)

	// start: 当前遍历的待切分位置
	var backTracking func(start int)
	backTracking = func(start int) {

		// 终止条件
		if start >= len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		// 单层处理逻辑
		for i := start; i < len(s); i++ {

			subStr := s[start : i+1]
			if isPalindrome(subStr) {
				path = append(path, subStr)

				backTracking(i + 1)       // 递归
				path = path[:len(path)-1] // 回溯
			}
		}
	}

	backTracking(0)

	return result

}

// isPalindrome 判断s是否是回文的
func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}

// restoreIpAddresses 93 - 复原IP地址
func restoreIpAddresses(s string) []string {

	result := make([]string, 0, 10)
	path := make([]string, 0, 10)

	// start 待切分位置
	var backTracking func(start int)
	backTracking = func(start int) {

		// 终止条件
		if start >= len(s) {
			if len(path) == 4 {
				ip := strings.Join(path, ".")
				result = append(result, ip)
			}
			return
		}

		// 单层处理逻辑
		for i := start; i < len(s); i++ {
			subStr := s[start : i+1]

			// 0开头的数字，例如01
			if len(subStr) > 1 && subStr[0] == 0 {
				continue
			}
			//超出范围
			subInt, _ := strconv.ParseInt(subStr, 10, 64)
			if subInt > 255 {
				break
			}

			path = append(path, subStr)

			backTracking(i + 1)       // 递归
			path = path[:len(path)-1] // 回溯

		}

	}

	backTracking(0)

	return result

}

// subSets 78 - 子集
func subSets(nums []int) [][]int {

	result := make([][]int, 0, 10)
	path := make([]int, 0, 10)

	var backTracking func(start int)
	backTracking = func(start int) {

		// 在每一层开始前收集结果
		// if len(path) > 0 {
		// 	temp := make([]int, 0, len(path))
		// 	copy(temp, path)
		// 	result = append(result, temp)
		// }

		// 空集也是子集
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)

		// 终止条件
		if start >= len(nums) {
			return
		}

		for i := start; i < len(nums); i++ {

			path = append(path, nums[i])
			backTracking(i + 1)       // 递归
			path = path[:len(path)-1] // 回溯

		}

	}

	backTracking(0)

	return result

}

// subsetsWithDup 90 - 子集II nums可能包含重复元素 - 排序+used数组
func subsetsWithDup(nums []int) [][]int {
	// 排序
	sort.Ints(nums)

	result := make([][]int, 0, 10)
	path := make([]int, 0, 10)
	var backTracking func(start int, used map[int]bool)
	backTracking = func(start int, used map[int]bool) {

		// 收集结果
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)

		// 终止条件
		if start >= len(nums) {
			return
		}

		// 单层处理逻辑
		for i := start; i < len(nums); i++ {

			// 去重 前一个相同的元素已经递推过，再递推就是重复
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			path = append(path, nums[i])
			used[i] = true

			backTracking(i+1, used) // 递归

			// 回溯
			path = path[:len(path)-1]
			used[i] = false

		}

	}

	backTracking(0, map[int]bool{})

	return result

}

// findSubsequences - 491 - 非递减子序列 - nums可能包含重复元素 - 排序+used数组
func findSubsequences(nums []int) [][]int {

	// 排序
	sort.Ints(nums)

	result := make([][]int, 0, 10)
	path := make([]int, 0, 10)

	var backTracking func(start int, used map[int]bool)
	backTracking = func(start int, used map[int]bool) {

		// 收集结果 多于2个元素
		if len(path) >= 2 {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
		}

		// 终止条件
		if start >= len(nums) {
			return
		}

		// 单层处理逻辑
		for i := start; i < len(nums); i++ {

			// 去重
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			// 	不满足递增条件
			if len(path) > 0 && nums[i] < path[len(path)-1] {
				continue
			}

			path = append(path, nums[i])
			used[i] = true
			backTracking(i+1, used) // 递归
			// 回溯
			path = path[:len(path)-1]
			used[i] = false
		}

	}

	backTracking(0, map[int]bool{})

	return result
}

// permute 4 - 全排列
// nums: 不包含重复数字
func permute(nums []int) [][]int {

	result := make([][]int, 0, 10)
	path := make([]int, 0, 10)

	var backTracking func(used map[int]bool)
	backTracking = func(used map[int]bool) {

		// 终止条件
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)

			return
		}

		for i := 0; i < len(nums); i++ {

			// 递归没用到的剩余集合
			if !used[i] {
				path = append(path, nums[i])
				used[i] = true

				// 递归
				backTracking(used)

				// 回溯
				path = path[:len(path)-1]
				used[i] = false
			}
		}

	}

	backTracking(map[int]bool{})

	return result

}

// permuteII 47 - 全排列II 
// nums: 可能包含重复的数字
func permuteII(nums []int) [][]int {

	// 排序
	sort.Ints(nums)

	result := make([][]int, 0, 10)
	path := make([]int, 0, 10)

	var backTracking func(used map[int]bool)
	backTracking = func(used map[int]bool) {

		// 终止条件
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)

			return
		}

		for i := 0; i < len(nums); i++ {

			// 这个条件说明这个位置上已经放过相同值的元素，重复了！
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			// 递归没用到的剩余集合
			if !used[i] {
				path = append(path, nums[i])
				used[i] = true

				// 递归
				backTracking(used)

				// 回溯
				path = path[:len(path)-1]
				used[i] = false
			}
		}

	}

	backTracking(map[int]bool{})

	return result
}
