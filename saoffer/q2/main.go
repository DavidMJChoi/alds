// 剑指 offer 面试题 2
// Given two strings representing binary numbers, calculate their sum.
// For example
// Input: 1 101
// Outout: 110

package main

import "fmt"

func main() {
	var a,b string

	fmt.Scan(&a,&b)

	var long, short string
	if len(a) > len(b) {
		long, short = a, b
	} else {
		long, short = b, a
	}

	for range len(long)-len(short) {
		short = "0" + short
	}

	res := ""
	c := byte(0)
	for i := len(long)-1; i >= 0; i-- {
		d1 := long[i] - 48
		d2 := short[i] - 48
		
		current := (d1+d2+c)%2
		c = byte((d1+d2+c)/2)

		res = string(current+48) + res
	}

	if c == 1 {
		res = "1" + res
	}

	fmt.Println(res)
}