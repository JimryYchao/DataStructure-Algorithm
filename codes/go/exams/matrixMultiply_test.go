package exams_test

import (
	"fmt"
	"testing"
)

// calculate C = C + A·B, if C = {0}, then C = A·B; Θ(n^3)
func MatrixMultiply(A, B, C [][]int, n int) {
	for i := range n { // n
		for j := range n { // n
			for k := range n { // n
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
}

func TestMatrixMultiply(t *testing.T) {
	A := [][]int{
		{1, 2},
		{3, 4},
	}
	B := [][]int{
		{4, 3},
		{2, 1},
	}
	C := make([][]int, 2)
	for i := range 2 {
		C[i] = make([]int, 2)
	}
	MatrixMultiply(A, B, C, 2)

	fmt.Println(C)

	A = [][]int{{1}}
	B = [][]int{{3}}
	C = [][]int{{0}}
	MatrixMultiply(A, B, C, 1)
	fmt.Println(C[0])
}

func SquareMatrixMultiply_Recursive(A, B, C [][]int, n int) { // T()
	if n <= 1 {
		C[0][0] += A[0][0] * B[0][0]
		return
	}
	// partition A,B, C into n/2 x n/2 sub-matrices:
	hn := n / 2
	A11, A12, A21, A22 := partM(A, hn, hn, n)
	B11, B12, B21, B22 := partM(B, hn, hn, n)
	C11, C12, C21, C22 := partM(C, hn, hn, n)

	SquareMatrixMultiply_Recursive(A11, B11, C11, hn)
	SquareMatrixMultiply_Recursive(A11, B12, C12, hn)
	SquareMatrixMultiply_Recursive(A21, B11, C21, hn)
	SquareMatrixMultiply_Recursive(A21, B12, C22, hn)
	SquareMatrixMultiply_Recursive(A12, B21, C11, hn)
	SquareMatrixMultiply_Recursive(A12, B22, C12, hn)
	SquareMatrixMultiply_Recursive(A22, B21, C21, hn)
	SquareMatrixMultiply_Recursive(A22, B22, C22, hn)

}

func partM(M [][]int, i, j, n int) ([][]int, [][]int, [][]int, [][]int) {
	// sup
	m11, m12 := make([][]int, i), make([][]int, i)
	for r := range i {
		m11[r] = M[r][:j]
		m12[r] = M[r][j:n]
	}
	M = M[i:]
	// sub
	m21, m22 := make([][]int, n-i), make([][]int, n-i)
	for r := range n - i {
		m21[r] = M[r][:j]
		m22[r] = M[r][j:n]
	}
	return m11, m12, m21, m22
}

func TestSquareMatrixMultiply_Recursive(t *testing.T) {
	A := [][]int{
		{1, 2},
		{3, 4},
	}
	B := [][]int{
		{4, 3},
		{2, 1},
	}
	C := make([][]int, 2)
	for i := range 2 {
		C[i] = make([]int, 2)
	}
	SquareMatrixMultiply_Recursive(A, B, C, 2)

	fmt.Println(C)
}

func Test(t *testing.T) {
	A := [][]int{
		{1, 2, 3, 4},
		{99, 98, 95, 94},
		{9, 7, 1, -1},
		{0, 20, 80, 100},
	}
	B := [][]int{
		{99, 98, 95, 94},
		{1, 2, 5, 6},
		{11, 22, 55, 66},
		{0, 20, 40, 60},
	}

	C := make([][]int, 4)
	for i := range 4 {
		C[i] = make([]int, 4)
	}
	MatrixMultiply(A, B, C, 4)
	fmt.Println(C)

	C2 := make([][]int, 4)
	for i := range 4 {
		C2[i] = make([]int, 4)
	}
	SquareMatrixMultiply_Recursive(A, B, C2, 4)
	fmt.Println(C2)
}
