// LeetCode L122
// DP: dp[i] = dp[i-1] + (prices[i] - prices[i-1]), if price diff is positive
//     dp[i] = dp[i-1], if price diff is non-positive
// use two cyclic variables to save space.

package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {

	// maxP := make([]int, len(prices))

	// maxP[0] = 0

	dp1 := 0
	dp2 := 0
	if len(prices) > 1 && prices[1]-prices[0] > 0 {
		dp2 = prices[1] - prices[0]
	}

	i := 1
	for i < len(prices) {
		if prices[i]-prices[i-1] > 0 {
			dp2 = dp1 + (prices[i] - prices[i-1])
		} else {
			dp2 = dp1
		}
		dp1 = dp2

		// fmt.Println(dp1, dp2)

		i++
	}

	return dp1
}
