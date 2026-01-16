// LeetCode L190

package main

import "fmt"

func main() {
	fmt.Println(reverseBits(43261596))
	// fmt.Println(reverseBits(2147483644))
}

func reverseBits(n int) int {
	ret := 0
	for range 32 {
		ret = (ret << 1) | (n & 1)
		n >>= 1
	}
	return ret
}
func reverseBits1(n int) int {
	nn := uint32(n)
	ret := uint32(0b0)

	// fmt.Printf("%032b\n%032b\n", nn, ret)

	// check if the most significant bit is 0 or 1
	for range 32 {
		s := nn & 0b10000000000000000000000000000000
		nn <<= 1
		ret >>= 1
		if s != 0 {
			ret |= 0b10000000000000000000000000000000
		}
	}

	// fmt.Printf("%032b\n", ret)

	return int(ret)
}
