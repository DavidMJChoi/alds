// LeetCode L189

package main

import "fmt"

func main() {
	rotate3([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 8)
	// rotate3([]int{-1, -100, 3, 99}, 3)
}

// 1. Trivial, move to provide space. In-place, O(1) extra space, O(n^2) time (very very slow)
// to avoid unnecessary repetition, get actual k first
func rotate1(nums []int, k int) {
	n := len(nums)

	k %= len(nums)

	for range k {
		tmp := nums[n-1]
		i := n - 1
		for i > 0 {
			nums[i] = nums[i-1]
			i--
		}
		nums[0] = tmp
	}
}

// 2. slicing. O(n) extra space, O(n) time
func rotate2(nums []int, k int) {
	n := len(nums)
	k %= n

	front := make([]int, n-k)
	rear := make([]int, k)

	copy(front, nums[:n-k])
	copy(rear, nums[n-k:])

	i := 0
	for i < k {
		nums[i] = rear[i]
		i++
	}
	j := 0
	for i < n {
		nums[i] = front[j]
		i++
		j++
	}

	fmt.Println(nums)
}

// 3. cyclic placement WORKING ON THIS...
func rotate3(nums []int, k int) {
	n := len(nums)
	k %= n

	tmp := 0
	i := n - k
	j := 0

	// m := n
	// if n%k == 0 {
	// 	m = k
	// }
	for range n {
		tmp = nums[j]
		nums[j] = nums[i]
		// get right position for tmp
		if j >= n-k {
			j -= (n - k)
		} else {
			j += k
		}
		nums[i] = tmp

		fmt.Println(nums, tmp)
	}
}
