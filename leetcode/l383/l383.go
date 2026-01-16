// LeetCode L383

package main

import "fmt"

func main() {
	fmt.Println(canConstruct("a", "b"))
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	cnt := make([]int, 26)
	for i := range magazine {
		cnt[magazine[i]-97]++
	}
	for i := range ransomNote {
		cnt[ransomNote[i]-97]--
		if cnt[ransomNote[i]-97] < 0 {
			return false
		}
	}

	return true
}
func canConstruct2(ransomNote string, magazine string) bool {
	hm := make(map[byte]int)

	for i := range magazine {
		v := magazine[i]
		_, ok := hm[v]
		if ok {
			hm[v]++
		} else {
			hm[v] = 1
		}
	}

	for i := range ransomNote {
		v := ransomNote[i]
		_, ok := hm[v]
		if !ok {
			return false
		} else {
			hm[v]--
			if hm[v] < 0 {
				return false
			}
		}
	}

	return true

}
func canConstruct3(ransomNote string, magazine string) bool {
	r := make(map[byte]int)
	m := make(map[byte]int)

	for i := range magazine {
		_, ok := m[magazine[i]]
		if ok {
			m[magazine[i]]++
		} else {
			m[magazine[i]] = 1
		}
	}
	for i := range ransomNote {
		_, ok := r[ransomNote[i]]
		if ok {
			r[ransomNote[i]]++
		} else {
			r[ransomNote[i]] = 1
		}
	}

	// fmt.Println(r)
	// fmt.Println(m)

	for k, v := range r {
		a, ok := m[k]
		if !ok {
			return false
		} else {
			if a < v {
				return false
			}
		}
	}

	return true
}
