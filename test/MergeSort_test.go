package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func MergeSort(A []int, p, r int) {
	if p >= r {
		return
	}
	q := (p + r) / 2
	MergeSort(A, p, q)
	MergeSort(A, q+1, r)
	merge(A, p, q, r)
}

func merge(A []int, p, q, r int) []int {
	AL, nl := make([]int, q+1-p), q+1-p // len of A[p:q+1]
	copy(AL, A[p:q+1])
	AR, nr := make([]int, r-q), r-q // len of A[q+1:r]
	copy(AR, A[q+1:r+1])
	i, j, k := 0, 0, p
	for i < nl && j < nr {
		if AL[i] <= AR[j] {
			A[k] = AL[i]
			i++
		} else {
			A[k] = AR[j]
			j++
		}
		k++
	}
	if i < nl {
		copy(A[k:], AL[i:])
	} else {
		copy(A[k:], AR[j:])
	}
	return A
}

func TestMergeSort(t *testing.T) {
	arr := []int{1, 2, 4, 3, 5}
	MergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func MergeInsertSort(A []int, p, r int, k int) {
	if r-p <= k {
		InsertSort(A[p:r+1], r+1-p)
		return
	}
	q := (p + r) / 2
	MergeInsertSort(A, p, q, k)
	MergeInsertSort(A, q+1, r, k)
	merge(A, p, q, r)
}

func TestMergeInsertSort(t *testing.T) {
	ra := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	arr := ra.Perm(50)
	MergeInsertSort(arr, 0, 49, 10)
	fmt.Println(arr)
}
