// LeetCode L763

package main

import "fmt"

func main() {
	fmt.Println(partitionLabels("ababcc"))
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("eccbbbbdec"))
}

func partitionLabels(s string) []int {
	n := len(s)

	maxReach := make(map[byte]int)
	// first pass: find each character's max reach
	for i, c := range s {
		maxReach[byte(c)] = i
	}

	// fmt.Println(maxReach)

	ret := []int{}

	i := 0
	for i < n {
		j := i
		currReach := maxReach[s[i]]
		for j <= currReach {
			if maxReach[s[j]] > currReach {
				currReach = maxReach[s[j]]
			}
			j++
		}
		ret = append(ret, j-i)
		i = j
	}

	return ret
}

func partitionLabels1(s string) []int {
	n := len(s)

	maxReach := make([]int, n)
	// first pass: find each character's max reach
	for i := range n {
		j := n - 1
		for j > i && s[j] != s[i] {
			j--
		}
		maxReach[i] = j
	}

	// fmt.Println(maxReach)

	ret := []int{}

	i := 0
	for i < n {
		j := i
		currReach := maxReach[i]
		for j <= currReach {
			if maxReach[j] > currReach {
				currReach = maxReach[j]
			}
			j++
		}
		ret = append(ret, j-i)
		i = j
	}

	return ret
}
