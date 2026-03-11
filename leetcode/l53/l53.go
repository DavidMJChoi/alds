// LeetCode L53
// check Notion for details and explanation

package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	curSum := nums[0]

	for _, num := range nums[1:] {
		curSum = max(num, curSum+num)
		maxSum = max(curSum, maxSum)
	}

	return maxSum
}
