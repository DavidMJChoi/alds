// LeetCode L26

package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	i := 0
	k := 1

	for i < len(nums) {
		j := i
		for j < len(nums) {
			if nums[j] != nums[i] {
				break
			}
			j++
		}
		fmt.Println(nums, i, j, k)
		if j < len(nums) {
			nums[k] = nums[j]
		}
		k++
		i = j
	}

	// fmt.Println(nums)

	return k - 1
}
