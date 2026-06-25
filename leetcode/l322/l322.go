// LeetCode L322

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{2}, 0))
	fmt.Println(coinChange([]int{2}, 1))
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for _, coin := range coins {
		if coin <= amount {
			dp[coin] = 1
		}
	}
	for i := 1; i <= amount; i++ {
		minn := math.MaxInt32
		for _, coin := range coins {
			if i-coin >= 0 {
				minn = min(dp[i-coin], minn)
			}
		}
		dp[i] = minn + 1
	}
	if dp[amount] == math.MaxInt32+1 {
		return -1
	}
	return dp[amount]
}
