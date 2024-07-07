package goA

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
