// LeetCode L27

package main

import "fmt"

func main() {
	nums1 := []int{3, 2, 2, 3}
	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}

	fmt.Println(removeElement(nums1, 3), nums1)
	fmt.Println(removeElement(nums2, 2), nums2)
}

func removeElement(nums []int, val int) int {
	i := 0
	k := 0
	for i < len(nums) {
		if nums[i] != val {
			nums[k] = nums[i]
			// if i > k {
			// 	nums[i] = -1
			// }
			k++
		}
		i++
	}
	// println(i, k)

	return len(nums) - k
}
