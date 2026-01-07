package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	res := []int{0}

	for i := 1; i <= n; i++ {
		res = append(res, res[i >> 1] + (i&1))
	}
	

	fmt.Println(res)
}