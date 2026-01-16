// LeetCode L67

package main

import "fmt"

func main() {
	fmt.Println(addBinary("11", "1"))
}

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		return addBinary(b, a)
	}
	// assume a is the longer one

	aa := make([]byte, len(a)+1)
	bb := make([]byte, len(a)+1)
	i := len(a) - 1
	j := len(aa) - 1
	for i >= 0 {
		aa[j] = a[i]
		i--
		j--
	}
	i = len(b) - 1
	j = len(bb) - 1
	for i >= 0 {
		bb[j] = b[i]
		i--
		j--
	}

	// fmt.Println(aa, bb)

	i = len(aa) - 1
	carry := byte(0)
	// println(carry)
	// fmt.Println(i)
	for i >= 1 {
		x := aa[i]
		y := bb[i]

		if x >= 48 {
			x -= 48
		}
		if y >= 48 {
			y -= 48
		}

		current := (x + y + carry) % 2
		carry = (x + y + carry) / 2

		// fmt.Println(current, carry)

		aa[i] = current + 48

		i--
	}

	if carry == 1 {
		aa[0] = '1'
		return string(aa)
	} else {
		return string(aa[1:])
	}

}
