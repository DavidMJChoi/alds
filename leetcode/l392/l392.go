// LeetCode L392
// Trivial

package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	i, j := 0, 0

	for i < len(s) && j < len(t) {
		for j < len(t) {
			if t[j] == s[i] {
				i++
				j++
				break
			}
			j++
		}
	}

	// fmt.Println(i, j)

	return i == len(s)
}
