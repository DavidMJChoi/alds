// LeetCode L35

package main

import "fmt"

func main() {
	for i := range 8 {
		fmt.Println(i, searchInsert([]int{1, 3, 5, 6}, i))
	}

}

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	m := (l + r) / 2

	for l <= r {
		m = (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return r + 1
}
