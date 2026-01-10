// LeetCode L5
// Trivial: easy. start from analyzing the longest possible subtring (the original string itself)
// Greedily find the longest possible palindrome

package main

import "fmt"

func main() {
	s1 := "babad"
	s2 := "cbbd"

	fmt.Println(longestPalindrome(s1))
	fmt.Println(longestPalindrome(s2))
}

func longestPalindrome(s string) string {
	subStringLength := len(s)
	for subStringLength > 1 {
		i := 0
		for i+subStringLength <= len(s) {
			currentSub := s[i : i+subStringLength]
			if isPalindrome(currentSub) {
				return currentSub
			}
			i++
		}
		subStringLength--
	}
	return s[:1]
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
