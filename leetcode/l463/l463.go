// LeetCode L463

package main

import "fmt"

func main() {
	fmt.Println(islandPerimeter([][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}))
}

func islandPerimeter(grid [][]int) int {
	i := 0

	perimeter := 0
	for i < len(grid) {
		j := 0
		for j < len(grid[i]) {
			if grid[i][j] == 1 {
				perimeter += 4
				if i-1 >= 0 && grid[i-1][j] == 1 {
					perimeter -= 2
				}
				if j-1 >= 0 && grid[i][j-1] == 1 {
					perimeter -= 2
				}
			}
			j++
			// fmt.Println(i, j)
			// fmt.Println(perimeter)
		}
		i++
		// fmt.Println(i, j)
	}

	return perimeter
}
