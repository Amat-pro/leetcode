package sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {

	arr := []int{8, 9, 10, 2, 5, 7, 8}
	selectionSort(arr)
	fmt.Println("selectionSort result: ", arr)

} 

func TestBubbleSort(t *testing.T) {
	arr := []int{8, 9, 10, 2, 5, 7, 8}
	bubbleSort(arr)
	fmt.Println("bubbleSort result: ", arr)
}

func TestQuickSort(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	quickSort(arr)
	fmt.Println("quickSort result: ", arr)
}