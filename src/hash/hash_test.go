package hash

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	fmt.Println("is anagram: ", isAnagram("anagram", "nagaram"))
}

func TestCommonChars(t *testing.T) {
	fmt.Println("common chars: ", commonChars([]string{"bella", "label", "roller"})) // e,l,l
}

func TestIntersection(t *testing.T) {
	fmt.Println("intersection: ", intersection([]int{1, 2, 2, 1}, []int{2, 2})) // 2
}

func TestIsHappy(t *testing.T) {
	fmt.Println("is happy: ", isHappy(19)) // true
	fmt.Println("is happy: ", isHappy(2)) // false
}
