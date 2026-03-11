// LeetCode L56

package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(merge([][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}))
	fmt.Println(merge([][]int{
		[]int{1, 4},
		[]int{4, 7},
	}))
	fmt.Println(merge([][]int{
		[]int{1, 4},
		[]int{0, 0},
	}))
	fmt.Println(merge([][]int{
		[]int{1, 4},
		[]int{2, 3},
	}))
}

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a []int, b []int) int {
		if a[0] < b[0] {
			return -1
		} else if a[0] == b[0] {
			return 0
		} else {
			return 1
		}
	})

	ret := [][]int{intervals[0]}

	i := 1
	j := 0
	for i < len(intervals) {
		cur := intervals[i]
		prev := ret[j]
		if cur[0] <= prev[1] && cur[1] >= prev[1] {
			prev[1] = cur[1]
		} else if cur[0] > prev[1] {
			ret = append(ret, cur)
			j++
		}
		i++
	}

	return ret
}
