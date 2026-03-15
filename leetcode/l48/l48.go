// LeetCode L48

package main

import "fmt"

func main() {
	rotate([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
}

func rotate(matrix [][]int) {
	n := len(matrix)
	// transpose, then reverse

	// reverse {
	for i := range n / 2 {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}

	// transpose
	for i := range n {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for _, row := range matrix {
		for _, ele := range row {
			fmt.Printf("%v ", ele)
		}
		fmt.Println()
	}
}
