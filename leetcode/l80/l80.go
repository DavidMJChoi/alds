// LeetCode L80

package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}

func removeDuplicates(nums []int) int {
	i := 0
	k := 0

	for i < len(nums) {
		j := i + 1
		twice := false
		for j < len(nums) {
			if nums[j] == nums[i] {
				twice = true
			}
			if nums[j] != nums[i] {
				break
			}
			j++
		}

		if twice {
			k += 2
		} else {
			k += 1
		}

		// eliminate dups exceeds the limit of 2
		if j < len(nums) {
			nums[k] = nums[j]
		}

		// go back 2 position to make sure the last twice-dup elements appear twice
		if twice {
			nums[k-2+1] = nums[k-2]
		}

		i = j
	}

	// fmt.Println(nums) ///

	return k
}
