// LeetCode L7
// I started to hate the type systme of Go...

package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(reverse(123))
	// fmt.Println(reverse(-123))
	// fmt.Println(reverse(120))
	fmt.Println(reverse(1563847412))

	// fmt.Println(reverse(-math.MaxInt32))
}

type int int32

func reverse(x int) int {
	digits := []int{}

	isNegative := false

	if x == -math.MaxInt32-1 {
		return 0
	}

	if x < 0 {
		isNegative = true
		x = -x
		// PROBLEM: -2^31
	}

	for x != 0 {
		digits = append(digits, x%10)
		x /= 10
	}

	// remove leading 0s
	start := 0
	for start < len(digits) {
		if digits[start] != 0 {
			break
		}
		start++
	}
	digits = digits[start:]

	// dealing with overflow
	// len(digits) == 10 and ...
	var result int = 0
	var base int = 1
	for i := len(digits) - 1; i >= 0; i-- {
		// dealing with potential overflow
		if len(digits) == 10 && i == 0 {
			var a int = digits[0]
			if a > 2 {
				return 0
			} else {
				result += digits[i] * base

				if result < 0 {
					return 0
				}
			}
		} else {
			result += digits[i] * base
			base *= 10
		}
	}

	if isNegative {
		result = -result
	}

	fmt.Println(digits, isNegative, result) ////////////////

	return result
}
