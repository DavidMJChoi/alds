// LeetCode L74

package main

import "fmt"

func main() {
	m1 := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	// m2 := [][]int{
	// {1},
	// }

	fmt.Println(searchMatrix(m1, 16))
	// fmt.Println(searchMatrix(m2, 0))
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	// binary search on column 1
	t := 0
	b := m - 1
	mid := (t + b) / 2

	for t <= b {
		mid = (t + b) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] > target {
			b = mid - 1
		} else {
			t = mid + 1
		}
	}

	// if not found in first column, find in bth row
	l := 0
	r := n - 1
	if b == -1 {
		b = 0
	}
	for l <= r {
		mid = (l + r) / 2
		if matrix[b][mid] == target {
			return true
		} else if matrix[b][mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}
