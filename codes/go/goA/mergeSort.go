package goA

func merge(A []int, l, m, r int) {
	nL, nR := m+1-l, r-m
	tmp := make([]int, r-l+1)
	// copy A[l:r] into tmp; and L = tmp[0:nL], R = tmp[nL:r+1]
	for i := 0; i < r-l+1; i++ {
		tmp[i] = A[l+i]
	}
	// merge
	iL, iR, k := 0, 0, l
	for iL < nL && iR < nR {
		if tmp[iL] < tmp[nL+iR] {
			A[k] = tmp[iL]
			iL++
		} else {
			A[k] = tmp[nL+iR]
			iR++
		}
		k++
	}
	// copy reminder
	for iL < nL {
		A[k] = tmp[iL]
		k++
		iL++
	}
	for iR < nR {
		A[k] = tmp[nL+iR]
		k++
		iR++
	}
}

func MergeSort(A []int, p, r int) {
	if p >= r {
		return
	}
	q := (p + r) / 2
	MergeSort(A, p, q)
	MergeSort(A, q+1, r)
	merge(A, p, q, r)
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
