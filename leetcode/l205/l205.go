// LeetCode L205
// 我是傻逼。
// 我确实是傻逼。
// Since isomorphism is order-preserving, we only need to check sequentially.
// check twice: 1. if s maps to t; 2. if s maps to s

package main

import (
	"fmt"
)

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("f11", "b23"))
	fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("bbbaaaba", "aaabbbba"))
	fmt.Println(isIsomorphic("badc", "baba"))
}

func isIsomorphic(s string, t string) bool {
	m := make(map[byte]byte)

	i := 0
	for i < len(s) {
		value, existed := m[s[i]]
		if !existed {
			m[s[i]] = t[i]
		} else {
			if value != t[i] {
				return false
			}
		}
		i++
	}

	n := make(map[byte]byte)
	i = 0
	for i < len(t) {
		value, existed := n[t[i]]
		if !existed {
			n[t[i]] = s[i]
		} else {
			if value != s[i] {
				return false
			}
		}
		i++
	}

	return true
}
