// LeetCode L228
// trivial...

package main

import "fmt"

func main() {
	s1 := []int{0, 1, 2, 4, 5, 7}

	fmt.Println(summaryRanges(s1))
	fmt.Println(len(summaryRanges(s1)))
}

func summaryRanges(nums []int) []string {

	ret := make([]string, 0)

	i := 0
	for i < len(nums) {
		j := i + 1
		for j < len(nums) && nums[j]-1 == nums[j-1] {
			j++
		}
		if i != j-1 {
			s := fmt.Sprintf("%d->%d", nums[i], nums[j-1])
			ret = append(ret, s)
		} else {
			ret = append(ret, fmt.Sprintf("%d", nums[i]))
		}

		// fmt.Println(i, j)
		i = j
	}

	return ret
}
