// LeetCode L121
// Brute-force is trivial. Not gonna discuss here.

// O(n) time can be achived by keeping track of the max price so far starting from the very right hand side (end of array)
// After that, maxRight[i] is the highest price seen after day i
// subtracting maxRight[i] by prices[i] and find the maximum value

// Note that we don't really need to store the maxRight for all day i

package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	// maxRight := make([]int, len(prices))

	currentMax := 0
	maxP := 0
	i := len(prices) - 1
	for i >= 0 {
		if prices[i] > currentMax {
			currentMax = prices[i]
		}
		// maxRight[i] = currentMax

		if currentMax-prices[i] > maxP {
			maxP = currentMax - prices[i]
		}

		i--
	}

	return maxP
}
