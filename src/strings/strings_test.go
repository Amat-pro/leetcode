package strings

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	s := []byte("hello")
	reverseString(s)
	fmt.Println("reverseString, result is: ", string(s))
}

func TestReverseStringII(t *testing.T) {
	fmt.Println("reverseString II, result is: ", reverseStr("abcdefg", 2)) // bacdfeg
	fmt.Println("reverseString II, result is: ", reverseStr("abcd", 2)) // bacd
}

func TestPathEncryption(t *testing.T) {
	fmt.Println("path encryption - 替换空格, result is: ", pathEncryption("We are happy."))
}

func TestReverseWords(t *testing.T) {
	fmt.Println("reverse words, result is: ", reverseWords(" the sky is blue "))
}

func TestReverseLeftWords(t *testing.T) {
	fmt.Println("reverse left words, result is: ", reverseLeftWords("abcdefg", 2)) // cdefgab
}

