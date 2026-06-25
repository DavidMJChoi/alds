// LeetCode L287

package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
}
func findDuplicate(nums []int) int {
	slow := 0
	fast := 0

	slow = nums[slow]
	fast = nums[fast]
	fast = nums[fast]

	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
		fast = nums[fast]
	}

	p := 0
	for p != slow {
		p = nums[p]
		slow = nums[slow]
	}

	return slow
}
