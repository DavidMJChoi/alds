// LeetCode L8

package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(myAtoi("42"))
	// fmt.Println(myAtoi("   -042"))
	// fmt.Println(myAtoi("1337c0d3"))
	// fmt.Println(myAtoi("0-1"))
	// fmt.Println(myAtoi("words and 987"))
	// fmt.Println(myAtoi("2147483648"))
	// fmt.Println(myAtoi("-2147483649"))
	// fmt.Println(myAtoi("-91283472332"))
	// fmt.Println(myAtoi("-+12"))
	// fmt.Println(myAtoi(""))
	// fmt.Println(myAtoi(" "))
	fmt.Println(myAtoi("20000000000000000000"))
}

func myAtoi(s string) int {

	if len(s) == 0 {
		return 0
	}

	// leading spaces
	i := 0
	for i < len(s) {
		if s[i] != ' ' {
			break
		}
		i++
	}

	// sign
	isNegative := false
	if i < len(s) {
		if s[i] == '-' {
			isNegative = true
			i++
		} else if s[i] == '+' {
			i++
		}
	}

	// leading 0s
	for i < len(s) {
		if s[i] != '0' {
			break
		}
		i++
	}

	s = s[i:]

	digits := []int{}

	i = 0
	for i < len(s) {
		if !(s[i] >= 48 && s[i] <= 57) {
			break
		}

		// fmt.Println(s[i])

		digits = append(digits, int(s[i]-48))
		i++
	}

	// fmt.Println(s, digits) ///////

	var result int32 = 0
	var base int32 = 1
	for i := len(digits) - 1; i >= 0; i-- {
		// dealing with potential overflow
		if len(digits) > 10 {
			if isNegative {
				return -math.MaxInt32 - 1
			} else {
				return math.MaxInt32
			}
		}
		if len(digits) == 10 && i == 0 {
			var a int = digits[0]
			if a > 2 {
				if isNegative {
					return -math.MaxInt32 - 1
				} else {
					return math.MaxInt32
				}
			} else {
				result += int32(digits[i]) * base

				if result < 0 {
					if isNegative {
						return -math.MaxInt32 - 1
					} else {
						return math.MaxInt32
					}
				}
			}
		} else {
			result += int32(digits[i]) * base

			// fmt.Println(result, isNegative)

			base *= 10
		}
	}

	if isNegative {
		result = -result
	}

	//

	return int(result)
}
