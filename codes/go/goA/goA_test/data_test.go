package goA_test

import (
	. "goDSA/goA"
	"math/rand"
	"testing"
	"time"
)

var _len = 4

func BenchmarkInsertSort(b *testing.B) {
	ra := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	arr0 := ra.Perm(_len)
	arr := make([]int, _len)
	for range b.N {
		copy(arr, arr0)
		InsertSort(arr, _len)
	}
}

func BenchmarkInsertSort_Re(b *testing.B) {
	ra := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	arr0 := ra.Perm(_len)
	arr := make([]int, _len)
	for range b.N {
		copy(arr, arr0)
		InsertSort_Recursive(arr, _len)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	ra := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	arr0 := ra.Perm(_len)
	arr := make([]int, _len)
	for range b.N {
		copy(arr, arr0)
		BubbleSort(arr, _len)
	}
}

func BenchmarkMergeInsertSort(b *testing.B) {
	ra := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	arr0 := ra.Perm(_len)
	arr := make([]int, _len)
	for range b.N {
		copy(arr, arr0)
		MergeInsertSort(arr, 0, _len-1, 200)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	ra := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	arr0 := ra.Perm(_len)
	arr := make([]int, _len)
	for range b.N {
		copy(arr, arr0)
		MergeSort(arr, 0, _len-1)
	}
}
