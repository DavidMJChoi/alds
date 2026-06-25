// LeetCode L70

package main

import "fmt"

func main() {
	for i := range 10 {
		if i == 0 {
			continue
		}
		fmt.Println(climbStairs(i))
	}
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	dp1 := 1
	dp2 := 2
	for range n - 2 {
		dp1, dp2 = dp2, dp1+dp2
	}

	return dp2
}
