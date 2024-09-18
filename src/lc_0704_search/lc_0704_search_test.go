package lc_0704_search

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {

	numbers := []int{-1, 0, 3,5,9,12}
	fmt.Println("target 9, result is: ", search(numbers, 9))
	fmt.Println("target 2, result is: ", search(numbers, 2))

}