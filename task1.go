package main

import (
	"fmt"
)

func findMaxPairIndexes(A, B []int) (int, int) {
	n := len(A)
	max_sum := A[0] + B[0]
	i0 := 0
	j0 := 0

	cur_sum := A[0] + B[0]

	j := 1
	for ; j < n; j++ {
		cur_sum = A[i0] + B[j]
		if cur_sum > max_sum {
			max_sum = cur_sum
			j0 = j
		}

		if cur_sum < A[j]+B[j] {
			i0 = j
			cur_sum = A[j] + B[j]
		}
	}

	return i0, j0
}

func main() {
	A := []int{4, -8, 6}
	B := []int{-10, 3, 1}

	i, j := findMaxPairIndexes(A, B)
	fmt.Println(i, j)
}
