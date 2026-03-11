// LeetCode L1768

package main

import (
	"fmt"
)

func main() {
	fmt.Println(mergeAlternately("abc", "pqr"))
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcd", "pq"))
}

func mergeAlternately(word1 string, word2 string) string {
	i := 0
	j := 0

	merged := make([]byte, len(word1)+len(word2))

	k := 0
	for i < len(word1) && j < len(word2) {
		merged[k] = word1[i]
		i++
		k++
		merged[k] = word2[j]
		j++
		k++
	}

	for i < len(word1) {
		merged[k] = word1[i]
		i++
		k++
	}

	for j < len(word2) {
		merged[k] = word2[j]
		j++
		k++
	}

	return string(merged)

	// var builder strings.Builder
	// for i < len(word1) && j < len(word2) {
	// 	builder.WriteByte(word1[i])
	// 	builder.WriteByte(word2[j])
	// 	i++
	// 	j++
	// }

	// fmt.Println(i, j)

	// for i < len(word1) {
	// 	builder.WriteByte(word1[i])
	// 	i++
	// }

	// for j < len(word2) {
	// 	builder.WriteByte(word2[j])
	// 	j++
	// }

	// return builder.String()
}
