// LeetCode L54

package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}))
	fmt.Println(spiralOrder([][]int{
		{2, 5},
		{8, 4},
		{0, -1},
	}))
}

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	out := []int{}

	// always turn right

	// initial position
	i, j := 0, 0
	direction := 1
	count := 0
	for count < m*n {
		// push current
		if matrix[i][j] != -101 {
			out = append(out, matrix[i][j])
			matrix[i][j] = -101
			count++
		}

		// take step
		// direction
		// 1 right
		// 2 down
		// 3 left
		// 4 up
		switch direction {
		case 1:
			// should go down
			if j+1 >= n || matrix[i][j+1] == -101 {
				direction = 2
				continue
			}
			j++
		case 2:
			// should go left
			if i+1 >= m || matrix[i+1][j] == -101 {
				direction = 3
				continue
			}
			i++
		case 3:
			// should go up
			if j-1 < 0 || matrix[i][j-1] == -101 {
				direction = 4
				continue
			}
			j--
		case 4:
			// should go right
			if i-1 < 0 || matrix[i-1][j] == -101 {
				direction = 1
				continue
			}
			i--
		}
	}

	return out
}
