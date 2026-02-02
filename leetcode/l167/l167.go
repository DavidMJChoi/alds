// LeetCode L167
// Use two pointers. Very trivial.
// This is medium? Seriously?

package main

import "fmt"

func main() {
	s1 := []int{2, 7, 11, 15}

	fmt.Println(twoSum(s1, 9))

}

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		} else {
			if numbers[i]+numbers[j] > target {
				j--
			} else {
				i++
			}
		}
	}

	return []int{0, 0}
}
