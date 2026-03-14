// LeetCode L238
// key1: avoid unnecessary calculation by observing that out[i] = out[i-1] / nums[i] * nums[i-1]
// key2: if 0 appears in nums[i], then out[k] = 0 for all k != i;
//		 if 0 appears more than once, out[i] = 0 for all i.

package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
	fmt.Println(productExceptSelf([]int{4, 3, 2, 1, 2}))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prod := make([]int, n)

	// it is guaranteed that n >= 2
	prod[0] = 1
	prod[1] = nums[0]
	i := 2
	for i < n {
		prod[i] = prod[i-1] * nums[i-1]
		i++
	}

	prodRight := 1
	j := n - 2
	for j >= 0 {
		prodRight *= nums[j+1]
		prod[j] *= prodRight
		j--
	}

	return prod
}

func productExceptSelf2(nums []int) []int {
	n := len(nums)
	prodLeft := make([]int, n)
	prodRight := make([]int, n)

	// it is guaranteed that n >= 2
	prodLeft[0] = 1
	prodLeft[1] = nums[0]
	i := 2
	for i < n {
		prodLeft[i] = prodLeft[i-1] * nums[i-1]
		i++
	}

	prodRight[n-1] = 1
	prodRight[n-2] = nums[n-1]
	j := n - 3
	for j >= 0 {
		prodRight[j] = prodRight[j+1] * nums[j+1]
		j--
	}

	prod := make([]int, n)
	for i := range n {
		prod[i] = prodLeft[i] * prodRight[i]
	}

	return prod
}

func productExceptSelf1(nums []int) []int {
	out := make([]int, len(nums))

	out[0] = 1
	for i := 1; i < len(nums); i++ {
		out[0] *= nums[i]
	}

	// fmt.Println(out[0])

	for i := 1; i < len(nums); i++ {
		if nums[i] != 0 {
			out[i] = out[i-1] / nums[i] * nums[i-1]
		} else {
			out[i] = 1
			for j := 0; j < len(nums); j++ {
				if j != i {
					out[i] *= nums[j]
				}
			}
		}
	}

	return out
}
