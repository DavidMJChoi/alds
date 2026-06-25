// LeetCode L45

package main

func main() {
	// fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	// fmt.Println(jump([]int{2, 3, 0, 1, 4}))
	// fmt.Println(jump([]int{2, 1}))
	// fmt.Println(jump([]int{1, 1, 1, 1}))
	// fmt.Println(jump([]int{3, 2, 1}))
}

// func jump(nums []int) int {
// 	steps := make([]int, len(nums))

// 	steps[0] = 0

// 	i := 1
// 	step := 1
// 	maxReach := nums[0]
// 	nextMaxReach := maxReach

// 	for i < len(steps) {

// 		if i <= maxReach {
// 			// fmt.Println(maxReach)
// 			if i+nums[i] > nextMaxReach {
// 				nextMaxReach = i + nums[i]
// 			}
// 			steps[i] = step
// 			i++
// 		} else {
// 			i = maxReach + 1
// 			maxReach = nextMaxReach
// 			step++
// 		}
// 		// i++
// 	}

// 	// fmt.Println(steps)

// 	return steps[len(steps)-1]
// }

func jump(nums []int) int {
	n := len(nums)
	steps := make([]int, n)
	steps[0] = 0

	i := 0
	step := 1
	maxReach := nums[0]
	nextMaxReach := 0
	for i < n {
		j := i + 1
		for j < n && j <= maxReach {
			steps[j] = step
			if j+nums[j] > nextMaxReach {
				nextMaxReach = j + nums[j]
			}
			j++
		}
		if j == n { // reached the end
			break
		}
		i = maxReach
		maxReach = nextMaxReach
		step++
	}

	return steps[n-1]
}
