// LeetCode L42
// Trivial: level by level. Very easy, very inefficient (time)

package main

import (
	"fmt"
	"slices"
)

type Stack []int

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(item int) {
	*s = append(*s, item) // Use append to add to the end
}

// Pop removes and returns the top element of the stack.
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false // Indicate that the stack is empty
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index] // Slice to remove the top element
	// Note: For stacks of pointers, maps, or other reference types,
	// you might want to zero the popped element to aid garbage collection.
	return element, true
}

// Peek returns the top element of the stack without removing it.
func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(*s) - 1
	return (*s)[index], true
}

func main() {
	// s3 := []int{5, 4, 1, 2}

	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	// fmt.Println(trap([]int{4, 2, 0, 3, `2, 5}))
	// fmt.Println(trap(s3))
	// fmt.Println(trap([]int{9, 6, 8, 8, 5, 6, 3}))
	// fmt.Println(trap([]int{1, 2, 3, 2, 3, 2, 1}))
}

func trap(height []int) int {
	trapped := 0

	maxLeftHeight := 0
	leftMax := make([]int, len(height))
	for i := 0; i < len(height); i++ {
		if height[i] > maxLeftHeight {
			maxLeftHeight = height[i]
		}
		leftMax[i] = maxLeftHeight
	}
	fmt.Println(leftMax) ///

	maxRightHeight := 0
	rightMax := make([]int, len(height))
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > maxRightHeight {
			maxRightHeight = height[i]
		}
		rightMax[i] = maxRightHeight
	}
	fmt.Println(rightMax) ///

	for i := 0; i < len(height); i++ {
		trapped += min(leftMax[i], rightMax[i]) - height[i]
	}

	return trapped
}

func trapFailed2(height []int) int {

	trapped := 0

	// trimming
	start := 1
	for start < len(height) {
		if height[start] < height[start-1] {
			start--
			break
		}
		start++
	}
	end := len(height) - 1
	for end > 0 {
		if height[end-1] < height[end] {
			end++
			break
		}
		end--
	}
	height = height[start:end]

	// fmt.Println(height) ////////////////

	if len(height) <= 2 {
		return 0
	}

	var stkIncrease Stack = Stack{}
	var stkDecrease Stack = Stack{}
	i := 1
	for i < len(height) {
		if height[i] > height[i-1] {
			j := i + 1
			for j < len(height) {
				if height[j] <= height[j-1] {
					break
				}
				j++
			}
			i = j - 1
			stkIncrease.Push(i)
		}

		if height[i] < height[i-1] {
			stkDecrease.Push(i - 1)
			j := i + 1
			for j < len(height) {
				if height[j] >= height[j-1] {
					break
				}
				j++
			}
			i = j - 1
		}

		i++
	}

	// fmt.Println(stkDecrease) //////////////
	// fmt.Println(stkIncrease) //////////////

	// merge intervals
	var stk Stack = Stack{}
	for !stkIncrease.IsEmpty() {
		v, _ := stkIncrease.Pop()
		stk.Push(v)
		v, _ = stkDecrease.Pop()
		stk.Push(v)
	}

	fmt.Println(stk) /////////////

	// i, j := 0, 1
	// for i < len(stkIncrease) {
	// 	if height[stkIncrease[j]] > height[stkIncrease[i]]
	// }

	for !stkDecrease.IsEmpty() {

		i, _ := stkDecrease.Pop()
		j, _ := stkIncrease.Pop()
		base := min(height[i], height[j])

		i++
		for i < j {
			trapped += base - height[i]
			i++
		}
	}

	return trapped
}

func trapFailed(height []int) int {

	if len(height) == 1 {
		return 0
	}

	trapped := 0

	var stk Stack = Stack{}
	i, j := 0, 1
	for j < len(height) {
		if height[j]-height[i] < 0 {
			stk.Push(height[i] - height[j])
		} else {
			for !stk.IsEmpty() {
				v, _ := stk.Pop()
				if height[i] > height[j] {
					trapped += v - (height[i] - height[j])
				} else {
					trapped += v
				}
			}
			i = j
		}

		if len(height)-1 == j {
			for i != j { // find forward
				if height[j-1] < height[j] {
					break
				} else {
					stk.Pop()
					j--
				}
			}
			for !stk.IsEmpty() {
				v, _ := stk.Pop()

				// fmt.Print(v, " ", v-(height[i]-height[j]), "\n") /////////////////

				if height[i] > height[j] {
					if v-(height[i]-height[j]) > 0 {
						trapped += v - (height[i] - height[j])

					}
				} else {
					trapped += v
				}

				// fmt.Println(stk, trapped) ///////////////
			}
			i = j
			break
		}

		j++
		fmt.Print(trapped, " ")
	}

	return trapped
}

func trapTrivial(height []int) int {

	trapped := 0

	maxl := slices.Max(height)

	for l := 1; l <= maxl; l++ {
		// level := make([]int, len(height))

		// first := 0
		// last := len(height) - 1
		i := 0
		for i < len(height) {
			if height[i] >= l {
				// first = i
				break
			}
			i++
		}
		j := len(height) - 1
		for j >= 0 {
			if height[j] >= l {
				// last = j
				break
			}
			j--
		}
		for i <= j {
			if height[i] < l {
				trapped++
			}
			i++
		}

	}
	return trapped
}
