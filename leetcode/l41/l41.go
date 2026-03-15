// LeetCode L41

package main

import "fmt"

func main() {
	fmt.Println(firstMissingPositive([]int{1, 1}))
	// fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	// fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	// fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}

func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := range n {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	for i := range n {
		idx := nums[i] - 1
		if nums[i] < 0 {
			idx = -nums[i] - 1
		}
		if idx < n && nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}

	}

	res := 0
	for res < n {
		if nums[res] >= 0 {
			break
		}
		res++
	}

	return res + 1
}

func firstMissingPositive1(nums []int) int {
	set := make(map[int]struct{})

	maxNum := 0
	for _, num := range nums {
		if num > 0 {
			set[num] = struct{}{}
			if num > maxNum {
				maxNum = num
			}
		}
	}

	i := 1
	for i <= maxNum {
		_, ok := set[i]

		if !ok {
			return i
		}
		i++
	}

	return i
}
