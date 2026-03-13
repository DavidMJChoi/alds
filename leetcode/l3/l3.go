// LeetCode Problem 3, longest non-duplicate substring
// Trivial: O(n^2). too slow.
// try KMP?
// sliding window: O(n)

package main

import "fmt"

func main() {
	// s1 := "abcabcbb"
	// s2 := "bbbbb"
	// s3 := "pwwkew"

	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	// fmt.Println(lengthOfLongestSubstring(s2))
	// fmt.Println(lengthOfLongestSubstring(s3))

	fmt.Println(lengthOfLongestSubstring("abba"))
}

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	// sliding window [i,j]
	i := 0
	j := 0
	maxLength := 1
	m := make(map[byte]int)
	for j < len(s) {
		if idx, ok := m[s[j]]; ok { // already exists
			for i <= idx {
				delete(m, s[i])
				i++
			}
		}
		// most recent position where s[j] appears
		m[s[j]] = j

		if j-i+1 > maxLength {
			maxLength = j - i + 1
		}
		j++
	}

	return maxLength
}

func lengthOfLongestSubstring2(s string) int {

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
