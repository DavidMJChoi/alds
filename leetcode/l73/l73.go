// LeetCode L73

package main

import "fmt"

func main() {
	setZeroes([][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	})
	setZeroes([][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	})
}

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	isFirstRowSet := false
	isFirstColSet := false

	for i := range m {
		if matrix[i][0] == 0 {
			isFirstColSet = true
		}
	}
	for j := range n {
		if matrix[0][j] == 0 {
			isFirstRowSet = true
		}
	}

	// mark rows and cols
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// set marked columns
	for col := 1; col < n; col++ {
		if matrix[0][col] == 0 {
			for i := range m {
				matrix[i][col] = 0
			}
		}
	}
	// set marked rows
	for row := 1; row < m; row++ {
		if matrix[row][0] == 0 {
			for j := range n {
				matrix[row][j] = 0
			}
		}
	}

	// deal with first row and first column
	if isFirstColSet {
		for i := range m {
			matrix[i][0] = 0
		}
	}
	if isFirstRowSet {
		for j := range n {
			matrix[0][j] = 0
		}
	}

	for _, row := range matrix {
		for _, ele := range row {
			fmt.Printf("%v ", ele)
		}
		fmt.Println()
	}
}

func setZeroes2(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	// using O(m+n) extra space
	// use two hash tables to record the rows and columns to be set to 0
	rows := make(map[int]struct{})
	cols := make(map[int]struct{})

	for i, row := range matrix {
		for j, ele := range row {
			if ele == 0 {
				rows[i] = struct{}{}
				cols[j] = struct{}{}
			}
		}
	}

	// set rows to 0
	for k := range rows {
		for i := range n {
			matrix[k][i] = 0
		}
	}
	// set cols to 0
	for k := range cols {
		for i := range m {
			matrix[i][k] = 0
		}
	}

	for _, row := range matrix {
		for _, ele := range row {
			fmt.Printf("%v ", ele)
		}
		fmt.Println()
	}
}
