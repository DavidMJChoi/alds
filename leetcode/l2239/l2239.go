// LeetCode L2239

package main

import "fmt"

func main() {
	fmt.Println(findClosestNumber([]int{-4, -2, 1, 4, 8}))
	fmt.Println(findClosestNumber([]int{2, -1, 1}))
}

func findClosestNumber(nums []int) int {
	minIndex := 0
	minDistance := nums[0]
	if minDistance < 0 {
		minDistance = -minDistance
	}

	i := 1
	for i < len(nums) {
		num := nums[i]
		curDistance := num
		if curDistance < 0 {
			curDistance = -curDistance
		}

		if curDistance < minDistance {
			minDistance = curDistance
			minIndex = i
		}

		if curDistance == minDistance {
			minDistance = curDistance
			if num > nums[minIndex] {
				minIndex = i
			}
		}

		i++
	}

	return nums[minIndex]
}
