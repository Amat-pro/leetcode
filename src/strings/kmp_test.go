package strings

import (
	"fmt"
	"testing"
)

func TestKMPNext(t *testing.T) {
	s := "aabaaf"
	next := next(s)
	fmt.Println("next is: ", next)
}

func TestKMPStrStr(t *testing.T) {
	haystack := "sadbutsad"
	needle := "sad"
	fmt.Println("kmp strStr, result is: ", strStr(haystack, needle)) // 0
}

func TestRepeatedSubStringPattern(t *testing.T) {
	fmt.Println("kmp repeated sub string pattern, result is: ", repeatedSubStringPattern("abababab"))
}
