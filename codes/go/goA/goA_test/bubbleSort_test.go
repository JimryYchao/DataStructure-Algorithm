package goA_test

import (
	"fmt"
	"goDSA/goA"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 25, 4, 8, 8, 7, 98, 2, 6, 8, 5, 98, 8, 4, 8, 44, 25, 8}
	goA.BubbleSort(arr, len(arr))
	fmt.Println(arr)
}
