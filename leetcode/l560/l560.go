// LeetCode L560

package main

import (
	"fmt"
)

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
	fmt.Println(subarraySum([]int{1, 2, 1, 2, 1}, 3))
	fmt.Println(subarraySum([]int{1, -1, 0}, 0))
}

func subarraySum(nums []int, k int) int {
	numSubarray := 0

	i := 0
	for i < len(nums) {
		j := i
		curSum := 0
		for j < len(nums) {
			curSum += nums[j]
			if curSum == k {
				numSubarray++
			}
			j++
		}
		i++
	}

	return numSubarray
}
