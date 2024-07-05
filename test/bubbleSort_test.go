package test

import (
	"fmt"
	"testing"
)

func BubbleSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		for j := n - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 25, 4, 8, 8, 7, 98, 2, 6, 8, 5, 98, 8, 4, 8, 44, 25, 8}
	BubbleSort(arr, len(arr))
	fmt.Println(arr)
}
