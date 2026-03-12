// LeetCode L128

package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{1, 0, 1, 2}))
	fmt.Println(longestConsecutive([]int{0, 1, 2, 4, 8, 5, 6, 7, 9, 3, 55, 88, 77, 99, 999999999}))
	fmt.Println(longestConsecutive([]int{}))
}

func longestConsecutive(nums []int) int {

	m := make(map[int]bool)

	for _, num := range nums {
		m[num] = true
	}

	maxStreak := 0
	for num := range m {
		if !m[num-1] {
			curStreak := 1
			currentNum := num
			for m[currentNum+1] {
				curStreak++
				currentNum++
			}
			if curStreak > maxStreak {
				maxStreak = curStreak
			}
		}
	}

	return maxStreak
}
