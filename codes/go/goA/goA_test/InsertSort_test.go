package goA_test

import (
	"fmt"
	. "goDSA/goA"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{1, 25, 4, 8, 8, 7, 98, 2, 6, 8, 5, 98, 8, 4, 8, 44, 25, 8}
	InsertSort(arr, len(arr))
	fmt.Println(arr)
}

func TestInsertSort_Recursive(t *testing.T) {
	arr := []int{1, 25, 4, 8, 8, 7, 98, 2, 6, 8, 5, 98, 8, 4, 8, 44, 25, 8}
	InsertSort_Recursive(arr, len(arr))
	InsertSort(arr, len(arr))
	fmt.Println(arr)
}
