package test

func Horner(arr []int, n, x int) int {
	p := 0
	for i := n - 1; i >= 0; i-- {
		p = arr[i] + x*p
	}
	return p
}
