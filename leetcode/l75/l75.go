// LeetCode L75

package main

import "fmt"

func main() {
	// sortColors2([]int{1, 0})
	// sortColors2([]int{2, 1, 2, 1, 0})
	sortColors2([]int{1, 2, 0})
}

func sortColors(nums []int) {
	// i is the next position to insert 0
	// j is the next position to insert 1
	i, j, k := 0, 0, 0
	n := len(nums)

	for k < n {
		if nums[k] == 0 {
			nums[i], nums[k] = nums[k], nums[i]
			// nums[i] is already filled with 1 found by k
			// and that 1 is moved to position k
			// we need to evict them to position j
			if i < j {
				nums[k], nums[j] = nums[j], nums[k]
			}
			i++
			j++
		} else if nums[k] == 1 {
			nums[j], nums[k] = nums[k], nums[j]
			j++
		}
		k++
	}
}

func sortColors2(nums []int) {
	// p0 is the next position to fill with 0
	// p2 is the next position to fill with 2
	p0 := 0
	p2 := len(nums) - 1

	i := 0
	for i <= p2 {
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			p0++
		} else if nums[i] == 2 {
			nums[p2], nums[i] = nums[i], nums[p2]
			p2--
		}
		i++
	}

	fmt.Println(nums)
}
