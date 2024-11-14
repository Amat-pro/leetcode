package string

import (
	"fmt"
	"testing"
)

func Test_reverseString(t *testing.T) {
	s := []byte("hello")
	reverseString(s)
	fmt.Println("reverseString result is: ", string(s))
}

func Test_reverseStr(t *testing.T) {
	s := "abcdefg"
	k := 2
	fmt.Println("reverseStr result is: bacdfeg ", reverseStr(s, k))
}

func Test_reverseWords(t *testing.T) {
	s := " the sky is blue "
	fmt.Println("reverseWords result is: ", reverseWords(s))
}

func Test_strStr(t *testing.T) {
	haystack := "sadbutsad"
	needle := "sad"

	fmt.Println("strStr result is: 0 ", strStr(haystack, needle))
}

func Test_repeatedSubstringPattern(t *testing.T) {
	fmt.Println("repeatedSubstringPattern result is: true ", repeatedSubstringPattern("abcabcabcabc"))
}
