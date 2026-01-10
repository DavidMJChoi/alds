// LeetCode L12
// try []byte. may increase speed

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToRoman(3749)) // MMMDCCXLIX
	// fmt.Println(intToRoman(58))
	// fmt.Println(intToRoman(1994))
}

// 4 IV, 9 IX, 40 XL, 90 XC, 400 CD, 900 CM

func intToRoman(num int) string {
	numRoman := ""

	d := num % 10
	c := num / 10 % 10
	b := num / 100 % 10
	a := num / 1000

	switch d {
	case 4:
		numRoman = "IV" + numRoman
	case 9:
		numRoman = "IX" + numRoman
	default:
		if d > 4 {
			numRoman = "V" + strings.Repeat("I", d-5) + numRoman
		} else {
			numRoman = strings.Repeat("I", d) + numRoman
		}
	}

	switch c {
	case 4:
		numRoman = "XL" + numRoman
	case 9:
		numRoman = "XC" + numRoman
	default:
		if c > 4 {
			numRoman = "L" + strings.Repeat("X", c-5) + numRoman
		} else {
			numRoman = strings.Repeat("X", c) + numRoman
		}
	}

	switch b {
	case 4:
		numRoman = "CD" + numRoman
	case 9:
		numRoman = "CM" + numRoman
	default:
		if b > 4 {
			numRoman = "D" + strings.Repeat("C", b-5) + numRoman
		} else {
			numRoman = strings.Repeat("C", b) + numRoman
		}
	}

	numRoman = strings.Repeat("M", a) + numRoman

	fmt.Println(a, b, c, d)

	return numRoman
}
