package goA

func InsertSort(A []int, n int) {
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1
		for ; j >= 0 && A[j] > key; j-- {
			A[j+1] = A[j]
		}
		A[j+1] = key
	}
}

// recursive insert sort
func InsertSort_Recursive(A []int, n int) {
	if n > 1 {
		q := n - 1
		InsertSort_Recursive(A, q)
		insert(A, q)
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
