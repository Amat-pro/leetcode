package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"testing"
)

// ===================================================== 数组 =====================================================
// 二分查找
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	return _search(nums, 0, len(nums)-1, target)

}

// _search [start, end]
func _search(nums []int, start, end int, target int) int {

	if start == end {
		if nums[start] == target {
			return start
		}

		return -1
	}

	middle := (start + end) / 2
	if nums[middle] > target {
		// 向左
		return _search(nums, start, middle-1, target)
	} else if nums[middle] < target {
		// 向右
		return _search(nums, middle+1, end, target)
	} else {
		return middle
	}

}

// 移除元素
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[left] = nums[right]
			left++
			right++

			continue
		}

		right++
		continue
	}

	return left

}

// 有序数组的平方
func sortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	left := 0
	right := len(nums) - 1

	result := make([]int, len(nums))
	k := 0
	for left < right {
		temp1 := nums[left] * nums[left]
		temp2 := nums[right] * nums[right]

		if temp1 <= temp2 {
			result[k] = temp1

			left++
			k++

			continue
		}

		result[k] = temp2
		right--
		k++

	}

	return result
}

// 长度最小的子数组
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 滑动窗口
	left := 0
	right := len(nums) - 1

	minSubLength := 0
	sum := math.MaxInt

	for right < len(nums) {
		sum += nums[right]

		// 持续移动left
		for sum >= target {
			temp := right - left + 1
			if temp < minSubLength {
				minSubLength = temp
			}

			sum -= nums[left]
			left++
		}

		right++
	}

	if minSubLength == math.MaxInt {
		return -1
	}

	return minSubLength

}

// ===================================================== 链表 =====================================================
type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) Println() {
	cur := n
	for cur != nil {
		fmt.Print(" ", cur.Val)
		cur = cur.Next
	}
}

// 移除链表元素
func removeElements(head *ListNode, val int) *ListNode {
	virtualHead := &ListNode{}
	virtualHead.Next = head

	pre := virtualHead
	cur := virtualHead.Next

	for cur != nil {
		if cur.Val != val {
			pre = cur
			cur = cur.Next

			continue
		}

		next := cur.Next
		pre.Next = next
		cur.Next = nil // !

		cur = next
	}

	return virtualHead.Next
}

// 反转链表
func reverseList(head *ListNode) *ListNode {

	pre := head
	cur := head.Next

	pre.Next = nil // !

	for cur != nil {
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}

	return pre
}

// 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	virtualHead := &ListNode{}

	virtualHead.Next = head

	pre := virtualHead
	cur := head

	for cur != nil && cur.Next != nil {
		next := cur.Next
		nextNext := next.Next

		cur.Next = nil
		next.Next = nil
		pre.Next = next
		next.Next = cur
		cur.Next = nextNext

		pre = cur
		cur = nextNext
	}

	return virtualHead.Next
}

func Test_swapPairs(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n6 := &ListNode{Val: 6}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	head := swapPairs(n1)
	head.Println()
}

// 删除链表的倒数第N个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 双指针 fast先走n个  然后slow和fast一起移动
	virtualHead := &ListNode{}
	virtualHead.Next = head

	fast := virtualHead
	slow := virtualHead

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {

		fast = fast.Next
		slow = slow.Next

	}

	nextNext := slow.Next.Next
	next := slow.Next

	slow.Next = nextNext
	next.Next = nil

	return virtualHead.Next

}

// getIntersectionNode 02.07 - 链表相交
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := 0
	lenB := 0

	tempNode := headA
	for tempNode != nil {
		lenA++
		tempNode = tempNode.Next
	}
	tempNode = headB
	for tempNode != nil {
		lenB++
		tempNode = tempNode.Next
	}

	// headA为长链表
	if lenA < lenB {
		lenA, lenB = lenB, lenA
		headA, headB = headB, headA
	}

	n := lenA - lenB
	for i := 0; i < n; i++ {
		headA = headA.Next
	}

	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil

}

// detectCycle 142 - 环形链表2
func detectCycle(head *ListNode) *ListNode {

	// 	快慢指针
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			temp := head
			for temp != slow {
				temp = temp.Next
				slow = slow.Next
			}
			return slow
		}

	}

	return nil

}

// ===================================================== 哈希表 =====================================================
// 有效的字母异位词
func isAnagram(s string, t string) bool {
	hashRecord := make([]int, 26)

	for _, c := range s {
		hashRecord[c-'a']++
	}

	for _, c := range t {
		hashRecord[c-'a']++
	}

	for _, v := range hashRecord {
		if v != 0 {
			return false
		}
	}

	return true

}

// 查找共用字符
func commonChars(words []string) []string {

	charsArray := make([][]int, 0, len(words))
	for i := 0; i < len(words); i++ {
		temp := make([]int, 26)
		charsArray = append(charsArray, temp)
	}

	for i, word := range words {
		for c := range word {
			charsArray[i][c-'a']++
		}
	}

	result := make([]string, 0, 10)
	for i := 0; i < 26; i++ {
		minValue := math.MaxInt
		for _, item := range charsArray {
			if item[i] < minValue {
				minValue = item[i]
			}
		}

		if minValue <= 0 {
			continue
		}

		for count := 0; count < minValue; count++ {
			result = append(result, string(rune(i)+'a'))
		}
	}

	return result

}

// 两个数组的交集
func intersection(nums1 []int, nums2 []int) []int {
	hash := map[int]int{}
	for _, num := range nums1 {
		if _, ok := hash[num]; !ok {
			hash[num] = num
		}
	}

	result := make([]int, 0, 10)
	for _, num := range nums2 {
		num1, ok := hash[num]
		if ok {
			result = append(result, num1)
			delete(hash, num1) // 去重
		}
	}

	return result
}

// 快乐数
func isHappy(n int) bool {

	hash := map[int]int{}

	for {

		sum := _happy(n)
		if sum == 1 {
			return true
		}

		if _, ok := hash[sum]; ok {
			return false
		}

		hash[sum] = sum
		n = sum

	}

}

func _happy(n int) int {
	sum := 0

	for n != 0 {
		temp := n % 10
		sum += temp * temp

		n = n / 10
	}

	return sum

}

func twoSum(nums []int, target int) []int {
	hash := map[int]int{}

	for i, v := range nums {
		idx, ok := hash[target-v]
		if ok {
			return []int{idx, i}
		}

		hash[v] = i

	}

	return []int{}
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

// 赎金信
func canConstruct(ransomNote string, magazine string) bool {
	hash := map[rune]int{}

	for _, char := range magazine {
		hash[char]++
	}

	for _, v := range ransomNote {
		hash[v]--
		if hash[v] < 0 {
			return false
		}
	}

	return true
}

// ===================================================== 双指针 =====================================================
// 反转字符串
func reverseString(s []byte) {
	left := 0
	right := len(s) - 1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// 反转字符串里的单词
func reverseStr(s string, k int) string {
	b := []byte(s)
	for i := 0; i < len(b); i += 2 * k {
		if i+k < len(b) {
			reverseString(b[i : i+k])
		} else {
			reverseString(b[i:])
		}
	}

	return string(b)
}

// reverseWords 反转字符串中的单词
func reverseWords(s string) string {
	b := removeSpace([]byte(s))

	// 反转整个字符串
	reverseString(b)
	// 反转每个单词
	start := 0
	for end := 0; end <= len(b); end++ {
		if end == len(b) || b[end] == ' ' { // end到末尾了或者end指向空格，则[start,end)是需要反转的单词
			reverseString(b[start:end])
			start = end + 1
		}
	}

	return string(b)
}

func removeSpace(b []byte) []byte {
	// 删除前边空格
	// 删除中间多余空格
	// 删除尾部空格

	slow := 0
	fast := 0
	for fast < len(b) && b[fast] == ' ' { // 移除前边的空格
		fast++
	}

	for fast < len(b) {
		if fast-1 > 0 && b[fast] == b[fast-1] && b[fast] == ' ' { // 保留每个单词后面第一个空格
			fast++
			continue
		}
		b[slow] = b[fast]
		slow++
		fast++
	}

	// 移除尾部的空格
	// slow-1是末尾位置（slow指向的是下一个写的位置，可能是等于length）
	if b[slow-1] == ' ' {
		return b[:slow-1]
	}
	return b[:slow]

}

// 三数之和
func threeSum(nums []int) [][]int {
	result := make([][]int, 0, 10)

	// 排序
	sort.Slice(nums, func(x, y int) bool {
		return nums[x] < nums[y]
	})

	for i := 0; i < len(nums); i++ {
		// 对i去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 双指针
		left := i + 1
		right := len(nums) - 1

		for left < right {
			temp := nums[i] + nums[left] + nums[right]
			if temp > 0 {
				right--
			} else if temp < 0 {
				left++
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 对left去重
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				// 对right去重
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return result
}

// ===================================================== 栈与队列 =====================================================
// 有效的括号
func isValid(s string) bool {
	stack := make([]rune, 0, len(s))

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if c == ')' && top != '(' {
				return false
			}
			if c == ']' && top != '[' {
				return false
			}
			if c == '}' && top != '{' {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// 删除字符串中的所有相邻重复项
func removeDumplicates(s string) string {
	stack := make([]rune, 0, len(s))

	for _, c := range s {
		if len(stack) > 0 && stack[len(stack)-1] == c {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, c)
	}

	return string(stack)

}

// 逆波兰表达式求值
func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	for _, token := range tokens {
		switch token {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			tokenNum, _ := strconv.Atoi(token)
			stack = append(stack, tokenNum)
		}
	}

	return stack[0]
}

// 滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	queue := NewQueue()
	for i := 0; i < k; i++ {
		queue.push(nums[i])
	}

	result := make([]int, 0, 10)
	// [0, k-1]窗口最大值
	result = append(result, queue.getTop())

	popIndex := 0
	pushIndex := k
	for pushIndex <= len(nums)-1 {
		queue.push(nums[pushIndex])
		queue.pop(nums[popIndex])

		result = append(result, queue.getTop())

		popIndex++
		pushIndex++

	}

	return result

}

type Queue struct {
	q []int
}

func NewQueue() *Queue {
	return &Queue{
		q: make([]int, 0, 10),
	}
}

func (queue *Queue) push(num int) {
	if queue.len() == 0 {
		queue.q = append(queue.q, num)
		return
	}

	// 向前卷走每一个比num小的元素
	for queue.len() > 0 && num > queue.getBack() {
		queue.q = queue.q[:queue.len()-1]
	}

	queue.q = append(queue.q, num)
}

func (queue *Queue) pop(num int) {
	if queue.len() == 0 {
		return
	}

	if num != queue.getTop() {
		return
	}

	// 弹出的是最大元素
	queue.q = queue.q[1:]
}

func (queue *Queue) getTop() int {
	return queue.q[0]
}

func (queue *Queue) getBack() int {
	return queue.q[queue.len()-1]
}

func (queue *Queue) len() int {
	return len(queue.q)
}

// ===================================================== 单调栈 =====================================================
// 每日温度
func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0, len(temperatures))

	result := make([]int, len(temperatures))
	for i, temperature := range temperatures {
		if len(temperatures) == 0 {
			stack = append(stack, temperature)
			continue
		}

		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}

	return result
}

// 下一个更大元素I

// 下一个更大元素II

// 接雨水
