// LeetCode L6
// Trivial method: easy

package main

import "fmt"

func main() {
	s1 := "PAYPALISHIRING"

	fmt.Println(convert(s1, 2))
	fmt.Println()
	fmt.Println(convert(s1, 3))
	fmt.Println()
	fmt.Println(convert(s1, 4))
	fmt.Println()
	fmt.Println(convert(s1, 5))
	fmt.Println()
	fmt.Println(convert(s1, 7))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	i := 0
	newString := ""
	for i < numRows {
		j := i
		zagLevel := numRows - i - 1
		if zagLevel == numRows-1 || zagLevel == 0 {
			zagLevel = 0
		}
		// fmt.Println(zagLevel)

		for j < len(s) {
			newString += string(s[j])
			if zagLevel != 0 && j+zagLevel*2 < len(s) && j+zagLevel*2 >= 0 {
				newString += string(s[j+zagLevel*2])
			}
			j += numRows*2 - 2
		}
		i++
	}
	return newString
}

// PAHNAPLSIIGYIR
