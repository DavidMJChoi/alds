//

package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	digits := []int{}
	for x != 0 {
		digits = append(digits, x%10)
		x /= 10
	}

	i, j := 0, len(digits)-1
	for i <= j {
		if digits[i] != digits[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-10))
	fmt.Println(isPalindrome(10))

	fmt.Println(int(5) % 1.0)
}
