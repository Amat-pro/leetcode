package main

import (
	"fmt"
	"sync"
)

// 2、定义三个不同的goroutine，其中第一个goroutine固定打印"A"，第二个goroutine固定打印"B"，第三个goroutine固定打印"C"，
// 现有一个输入n，表示循环"ABC"字符串的数量，请按照期望输出指定的结果。
// 例如：
// 输入：1
// 输出："ABC"
// 输入：3
// 输出："ABCABCABC"

func printABC(n int) {

	channelA := make(chan struct{}, 0)
	channelB := make(chan struct{}, 0)
	channelC := make(chan struct{}, 0)

	wg := sync.WaitGroup{}
	wg.Add(1)
	printAFn := func(c chan struct{}) {
		// defer wg.Done()
		for _ = range c {
			fmt.Print("A")
			channelB <- struct{}{}
		}
	}
	printBFn := func(c chan struct{}) {
		for _ = range c {
			fmt.Print("B")
			channelC <- struct{}{}
		}
	}
	printCFn := func(c chan struct{}, count int) { // 加一个计数器
		i := 0
		for _ = range c {
			fmt.Print("C")
			i++
			if count == i {
				break
			}
		}
		wg.Done()
	}

	go printAFn(channelA)
	go printBFn(channelB)
	go printCFn(channelC, n)

	for i := 0; i < n; i++ {
		channelA <- struct{}{}
	}

	// time.Sleep(3 *time.Second)
	wg.Wait()

}

// 动态规划
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
// 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
// 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

// 示例 1：
// 输入：text1 = "abcde", text2 = "ace"
// 输出：3
// 解释：最长公共子序列是 "ace" ，它的长度为 3 。
// 示例 2：
// 输入：text1 = "abc", text2 = "abc"
// 输出：3
// 解释：最长公共子序列是 "abc" ，它的长度为 3 。
// 示例 3：
// 输入：text1 = "abc", text2 = "def"
// 输出：0
// 解释：两个字符串没有公共子序列，返回 0 。
