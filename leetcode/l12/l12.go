// LeetCode L12
// try []byte. may increase speed

package main

import (
	"fmt"
)

func main() {
	fmt.Println(intToRoman(3749)) // MMMDCCXLIX
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
	fmt.Println(intToRoman(388))
	fmt.Println(intToRoman(3888))
}

// 4 IV, 9 IX, 40 XL, 90 XC, 400 CD, 900 CM

func intToRoman(num int) string {
	numRoman := make([]byte, 15)

	d := num % 10
	c := num / 10 % 10
	b := num / 100 % 10
	a := num / 1000

	i := 14

	switch d {
	case 4:
		numRoman[i] = 'V'
		numRoman[i-1] = 'I'
		i -= 2
		// numRoman = "IV" + numRoman
	case 9:
		numRoman[i] = 'X'
		numRoman[i-1] = 'I'
		i -= 2
		// numRoman = "IX" + numRoman
	default:
		if d > 4 {
			for range d - 5 {
				numRoman[i] = 'I'
				i--
			}
			numRoman[i] = 'V'
			i--
			// numRoman = "V" + strings.Repeat("I", d-5) + numRoman
		} else {
			for range d {
				numRoman[i] = 'I'
				i--
			}
			// numRoman = strings.Repeat("I", d) + numRoman
		}
	}
	switch c {
	case 4:
		numRoman[i] = 'L'
		numRoman[i-1] = 'X'
		i -= 2
		// numRoman = "IV" + numRoman
	case 9:
		numRoman[i] = 'C'
		numRoman[i-1] = 'X'
		i -= 2
		// numRoman = "IX" + numRoman
	default:
		if c > 4 {
			for range c - 5 {
				numRoman[i] = 'X'
				i--
			}
			numRoman[i] = 'L'
			i--
			// numRoman = "V" + strings.Repeat("I", d-5) + numRoman
		} else {
			for range c {
				numRoman[i] = 'X'
				i--
			}
			// numRoman = strings.Repeat("I", d) + numRoman
		}
	}

	// switch c {
	// case 4:
	// 	numRoman = "XL" + numRoman
	// case 9:
	// 	numRoman = "XC" + numRoman
	// default:
	// 	if c > 4 {
	// 		numRoman = "L" + strings.Repeat("X", c-5) + numRoman
	// 	} else {
	// 		numRoman = strings.Repeat("X", c) + numRoman
	// 	}
	// }

	switch b {
	case 4:
		numRoman[i] = 'D'
		numRoman[i-1] = 'C'
		i -= 2
		// numRoman = "IV" + numRoman
	case 9:
		numRoman[i] = 'M'
		numRoman[i-1] = 'C'
		i -= 2
		// numRoman = "IX" + numRoman
	default:
		if b > 4 {
			for range b - 5 {
				numRoman[i] = 'C'
				i--
			}
			numRoman[i] = 'D'
			i--
			// numRoman = "V" + strings.Repeat("I", d-5) + numRoman
		} else {
			for range b {
				numRoman[i] = 'C'
				i--
			}
			// numRoman = strings.Repeat("I", d) + numRoman
		}
	}

	// switch b {
	// case 4:
	// 	numRoman = "CD" + numRoman
	// case 9:
	// 	numRoman = "CM" + numRoman
	// default:
	// 	if b > 4 {
	// 		numRoman = "D" + strings.Repeat("C", b-5) + numRoman
	// 	} else {
	// 		numRoman = strings.Repeat("C", b) + numRoman
	// 	}
	// }

	for range a {
		numRoman[i] = 'M'
		i--
	}

	// numRoman = strings.Repeat("M", a) + numRoman

	// fmt.Println(a, b, c, d)

	return string(numRoman[i+1:])
}
