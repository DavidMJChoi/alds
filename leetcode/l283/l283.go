// LeetCode L283

package main

import "fmt"

func main() {
	s1 := []int{0, 0, 1}

	moveZeroes(s1)

	fmt.Println(s1)
}

func moveZeroes(nums []int) {
	i, j := 0, 0
	n := len(nums)
	for j < n {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
		j++
	}

	for i < j {
		nums[i] = 0
		i++
	}
}
