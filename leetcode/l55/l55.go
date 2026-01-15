// LeetCode L55

// dp. Whether or not reaching len-1 depends on whether dp[len-2] >= 1
// canJump2 uses in-place dp, it is much slower than canJump1
// WHY?
// 版本1：创建新数组
// dp := make([]int, len(nums))
// 这是分配连续内存，通常很快

// 版本2：原地修改
// 每次迭代都需要从内存读取 nums[i] 的原始值
// 如果数组很大，可能会有缓存不友好的访问模式

package main

import "fmt"

func main() {

	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{0}))
	fmt.Println(canJump([]int{0, 1}))
	fmt.Println(canJump([]int{2, 0, 0}))
	fmt.Println(canJump([]int{0, 2, 3}))
}

func canJump1(nums []int) bool {
	dp := make([]int, len(nums))

	if len(nums) == 1 {
		return true
	}

	dp[0] = nums[0]

	i := 1
	for i < len(dp) {
		if dp[i-1] >= 1 {
			dp[i] = max(dp[i-1]-1, nums[i])
		}
		i++
	}

	return dp[len(dp)-2] >= 1
}

func canJump2(nums []int) bool {

	if len(nums) == 1 {
		return true
	}

	i := 1
	for i < len(nums) {
		if nums[i-1] >= 1 {
			nums[i] = max(nums[i-1]-1, nums[i])
		} else {
			nums[i] = 0
		}
		i++
	}

	fmt.Println(nums)

	return nums[len(nums)-2] >= 1
}

func canJump(nums []int) bool {

	if len(nums) == 1 {
		return true
	}

	dp1 := nums[0]
	dp2 := 0

	i := 1
	for i < len(nums)-1 {
		if dp1 >= 1 {
			dp2 = max(dp1-1, nums[i])
		} else {
			dp2 = 0
		}
		i++
		dp1 = dp2
	}

	// fmt.Println(dp1)

	return dp1 >= 1
}
