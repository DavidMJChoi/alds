// LeetCode L202

package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
	fmt.Println(isHappy(3))

	// fmt.Println(sumOfSquare(3))
	// fmt.Println(sumOfSquare(9))
	// fmt.Println(sumOfSquare(81))
	// fmt.Println(sumOfSquare(65))
	// fmt.Println(sumOfSquare(61))
	// fmt.Println(sumOfSquare(37))
	// fmt.Println(sumOfSquare(58))
	// fmt.Println(sumOfSquare(89))
}
func isHappy(n int) bool {

	hm := make(map[int]bool)

	sumOfSq := sumOfSquare(n)

	for sumOfSq != 1 {

		sumOfSq = sumOfSquare(sumOfSq)

		_, ok := hm[sumOfSq]
		if !ok {
			hm[sumOfSq] = true
		} else {
			return false
		}

	}

	return true
}

func sumOfSquare(n int) int {
	ret := 0
	for n != 0 {
		ret += (n % 10) * (n % 10)
		n /= 10
	}
	return ret
}
