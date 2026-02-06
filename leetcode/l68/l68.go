// LeetCode L80

package main

import (
	"bytes"
	"fmt"
)

func main() {
	// w1 := []string{"This", "is", "an", "example", "of", "text", "justification."}
	// w1 := []string{"enough", "to", "explain", "to"}
	// w2 := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	// w3 := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	w4 := []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"}

	// for _, line := range fullJustify(w1, 16) {
	// 	fmt.Println(line)
	// }
	// for _, line := range fullJustify(w2, 16) {
	// 	fmt.Println(line)
	// }
	// for _, line := range fullJustify(w3, 20) {
	// 	fmt.Println(line)
	// }

	// justified1 := fullJustify(w1, 16)
	// fmt.Println(len(justified1[0]), justified1)

	fmt.Println(fullJustify(w4, 16))
}

func fullJustify(words []string, maxWidth int) []string {
	out := make([]string, 0)

	i := 0
	for i < len(words) {

		currentLength := 0
		j := i
		for j < len(words) && currentLength < maxWidth {
			currentLength += len(words[j])
			currentLength += 1

			j++

			// fmt.Println(j, currentLength)
		}
		if currentLength > maxWidth+1 {
			j--
		}

		// last line: left justify
		if j == len(words) {
			buf := bytes.NewBuffer(make([]byte, maxWidth))
			buf.Reset()
			for _, word := range words[i:j] {
				buf.WriteString(word)
				buf.WriteByte(' ')
			}
			for buf.Len() < maxWidth {
				buf.WriteByte(' ')
			}
			out = append(out, buf.String())
		} else {
			// full justify
			buf := bytes.NewBuffer(make([]byte, maxWidth))
			buf.Reset()
			textLength := 0
			for _, word := range words[i:j] {
				textLength += len(word)
			}
			spaceLength := maxWidth - textLength
			spaces := make([]int, j-i)

			if len(spaces) == 1 {
				spaces[0] = spaceLength
			} else {
				originSpaceLength := spaceLength
				for n := range spaces[:len(spaces)-1] {
					spaces[n] = originSpaceLength / (len(spaces) - 1)
					spaceLength -= spaces[n]
				}
				m := 0
				for spaceLength > 0 {
					spaces[m] += 1
					spaceLength -= 1
					m++
				}
			}
			for k, word := range words[i:j] {
				buf.WriteString(word)

				if j-i == 1 {
					for range spaces[0] {
						buf.WriteByte(' ')
					}
				}

				if k != j-i-1 {
					for range spaces[k] {
						buf.WriteByte(' ')
					}
				}
			}
			out = append(out, buf.String())
		}

		i = j
	}

	return out
}
