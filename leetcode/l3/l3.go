// LeetCode Problem 3, longest non-duplicate substring
// Trivial: O(n^2). too slow.
// try KMP?

package main

import "fmt"

func main() {
	s1 := "abcabcbb"
	s2 := "bbbbb"
	s3 := "pwwkew"

	fmt.Println(lengthOfLongestSubstring(s1))
	fmt.Println(lengthOfLongestSubstring(s2))
	fmt.Println(lengthOfLongestSubstring(s3))
}

func lengthOfLongestSubstring(s string) int {

	maxLength := 0

	i := 0
	for i < len(s) {
		j := i + 1
		currentLength := 1
		currentChars := make(map[byte]int)

		currentChars[s[i]] = i

		for j < len(s) {
			position, ok := currentChars[s[j]]

			if ok {
				i = position
				break
			}

			currentChars[s[j]] = j
			currentLength += 1
			j++
		}
		if currentLength > maxLength {
			maxLength = currentLength
		}
		i++
	}

	return maxLength
}
