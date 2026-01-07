package main

import (
	"fmt"
	"log"
	"math"
)
func main() {

	var a int8 = 0
	var b int8 = 0
	
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		log.Fatalln("Fatal.")
	}

	if a == 0 {
		fmt.Println(0)
		return 
	}

	if a == math.MinInt8 && b == -1 {
		fmt.Println(math.MaxInt8)
		return
	}

	neg := false
	if a > 0 && b < 0 || a < 0 && b > 0 {
		neg = true
	}

	if a > 0 {
		a = -a
	}
	if b > 0 {
		b = -b
	}

	q := 0

	i := b
	j := 1

	bs := []int8{i}
	ls := []int{j}

	for i > a {
		i += i
		j += j
		bs = append(bs, i)
		ls = append(ls, j)
	}

	bs = bs[:len(bs)-1]
	ls = ls[:len(ls)-1]
	fmt.Println(bs)
	fmt.Println(ls)

	

	for i := len(bs)-1; i >= 0; i-- {
		a -= bs[i]
		q += ls[i]

		fmt.Println(a)
	}

	for a <= b {
		a -= b
		q++
	}

	if neg {
		q = -q
	}

	fmt.Println(q)
}