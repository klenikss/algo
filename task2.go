package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

func intersection(A, B []int) []int {
	intersect := make([]int, 0)

	k := binarySearch(A, B[len(B)-1])

	for i := range B {
		left := max(0, binarySearch(A[:k+1], B[i]))
		right := binarySearch(A[left:k+1], B[i])

		if right != -1 {
			intersect = append(intersect, B[i])
		}
	}

	return intersect
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	A := []int{1, 2, 3, 4, 5}
	B := []int{1, 3, 5}
	result := intersection(A, B)
	fmt.Println(result)
}