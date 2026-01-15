// LeetCode L274

package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(hIndex([]int{3, 3, 0, 6, 1, 5}))
	fmt.Println(hIndex([]int{1, 3, 1}))
	fmt.Println(hIndex([]int{100}))
	fmt.Println(hIndex([]int{0, 0, 3, 3}))
}

func hIndex(citations []int) int {
	slices.Sort(citations)
	// fmt.Println(citations)

	i := len(citations) - 1
	h := 0
	for i >= 0 {

		if h >= citations[i] {
			break
		}

		h++
		i--
	}

	return h
}
