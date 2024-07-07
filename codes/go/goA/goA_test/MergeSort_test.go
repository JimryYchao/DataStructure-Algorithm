package goA_test

import (
	"fmt"
	. "goDSA/goA"
	"math/rand"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	ra := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	arr := ra.Perm(50)
	MergeSort(arr, 25, 49)
	fmt.Println(arr)
}

func TestMergeInsertSort(t *testing.T) {
	ra := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	arr := ra.Perm(50)
	MergeInsertSort(arr, 25, 49, 10)
	fmt.Println(arr)
}
