package test

import (
	"fmt"
	"testing"
)

func InsertSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
}

func TestInsertSort(t *testing.T) {
	arr := []int{1, 25, 4, 8, 8, 7, 98, 2, 6, 8, 5, 98, 8, 4, 8, 44, 25, 8}
	InsertSort(arr, len(arr))
	fmt.Println(arr)
}

func InsertSort_Recursive(arr []int, n int) {
	if n > 1 {
		q := n - 1
		InsertSort_Recursive(arr, q)
		insert(arr, q)
	}
}
func insert(arr []int, r int) {
	j := r - 1
	key := arr[r]
	for ; j >= 0 && arr[j] > key; j-- {
		arr[j+1] = arr[j]
	}
	arr[j+1] = key
}
func TestInsertSort_Recursive(t *testing.T) {
	arr := []int{1, 25, 4, 8, 8, 7, 98, 2, 6, 8, 5, 98, 8, 4, 8, 44, 25, 8}
	InsertSort_Recursive(arr, len(arr))
	InsertSort(arr, len(arr))
	fmt.Println(arr)
}
