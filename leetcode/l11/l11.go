// LeetCode L11

package main

import "fmt"

func main() {
	s1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	s2 := []int{1, 1}

	fmt.Println(maxArea(s1))
	fmt.Println(maxArea(s2))
}

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1

	maxa := 0
	for i < j {
		h := min(height[i], height[j])

		a := h * (j - i)
		if a > maxa {
			maxa = a
		}

		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return maxa
}
