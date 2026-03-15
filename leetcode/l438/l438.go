// LeetCode L438

package main

import (
	"fmt"
)

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}

func findAnagrams(s string, p string) []int {
	ret := []int{}

	m := make(map[byte]struct{})
	for _, c := range []byte(p) {
		m[c] = struct{}{}
	}

	fmt.Println(m)

	i := 0
	for i <= len(s)-len(p) {
		// j := i

		n := make(map[byte]struct{})
		for _, c := range []byte(s[i : i+len(p)]) {
			n[c] = struct{}{}
		}

		flag := true
		for k := range n {
			if _, ok := m[k]; !ok {
				flag = false
				break
			}
		}
		for k := range m {
			if _, ok := n[k]; !ok {
				flag = false
				break
			}
		}
		if flag {
			ret = append(ret, i)
		}

		// for j < i+len(p) {
		// 	if _, ok := m[s[j]]; !ok {
		// 		i = j + 1
		// 		break
		// 	}
		// 	j++
		// 	if j == i+len(p) {
		// 		ret = append(ret, i)
		// 	}
		// }
		i++
	}

	return ret
}
