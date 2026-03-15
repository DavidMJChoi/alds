// LeetCode L240

package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 20))

	fmt.Println(searchMatrix([][]int{
		{-5},
	}, -5))
}

// Z search
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	i, j := 0, n-1

	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}

	return false
}

// binary search
func searchMatrix1(matrix [][]int, target int) bool {
	// m := len(matrix)
	n := len(matrix[0])

	m := binarySearchOnColumn(matrix, 0, target)
	// fmt.Println(m)

	// binary search on each row
	r := n - 1
	for i := range m {
		// fmt.Println(r)
		l := 0
		// r = n - 1
		for l <= r {
			m := (l + r) / 2
			if matrix[i][m] == target {
				return true
			} else if matrix[i][m] > target {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return false
}

func binarySearchOnColumn(matrix [][]int, col, target int) int {
	m, pos := len(matrix), len(matrix)

	l := 0
	r := m - 1
	for l <= r {
		m := (l + r) / 2
		if matrix[m][col] > target {
			pos = m
			r = m - 1
		} else if matrix[m][col] < target {
			l = m + 1
		} else {
			pos = m + 1
			return pos
		}
	}

	return pos
}
