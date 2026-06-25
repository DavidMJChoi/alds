// LeetCode L279

package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; i <= 16; i++ {
		fmt.Println(i, numSquares(i))
	}

	fmt.Println(numSquares(12))
}

func numSquares(n int) int {
	dp := make([]int, n+1)

	for i := 1; i < n+1; i++ {
		curMin := i
		k := 1
		for k*k < int(math.Sqrt(float64(i))) {
			k++
		}
		for k*k <= i {
			if dp[i-k*k] < curMin {
				curMin = dp[i-k*k] + 1
			}
			k++
		}
		dp[i] = curMin
	}
	return dp[n]
}

func numSquaresFailed(n int) int {
	// fmt.Println(n)
	if n == 0 {
		return 0
	}

	k := 1
	for k*k <= n {
		k++
	}
	k--
	return numSquares(n-k*k) + 1
}
