// LeetCode L198

package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{2, 1, 1, 2}))

}

func rob(nums []int) int {
	maxProfitRob := nums[0]
	maxProfitSKip := 0
	if len(nums) >= 2 {
		maxProfitRob = nums[1]
		maxProfitSKip = nums[0]
	}

	for i := 2; i < len(nums); i++ {
		maxProfitRob, maxProfitSKip = nums[i]+maxProfitSKip, max(maxProfitSKip, maxProfitRob)
	}

	return max(maxProfitRob, maxProfitSKip)
}
