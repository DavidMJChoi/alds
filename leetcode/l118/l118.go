// LeetCode L118

package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	triangle := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		row := []int{1}
		prevRow := triangle[i-1]
		for j := 0; j < len(prevRow)-1; j++ {
			row = append(row, prevRow[j]+prevRow[j+1])
		}
		row = append(row, 1)
		triangle = append(triangle, row)
	}
	return triangle
}
