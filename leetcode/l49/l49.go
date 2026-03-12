// LeetCode L49

package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

// count instead of sort... not really improved...
func groupAnagrams2(strs []string) [][]string {
	// use a hash map to record anagrams
	anagrams := make(map[[26]byte][]string)

	for _, s := range strs {
		// sort each string first.
		count := [26]byte{}
		for i := 0; i < len(s); i++ {
			count[s[i]-'a']++
		}

		// fmt.Println(ss)

		anagrams[count] = append(anagrams[count], s)

		// fmt.Println(anagrams)

	}
	ret := make([][]string, 0, len(anagrams))
	for _, v := range anagrams {
		ret = append(ret, v)
	}

	return ret
}

func groupAnagrams(strs []string) [][]string {
	// use a hash map to record anagrams
	anagrams := make(map[string][]string)

	for _, s := range strs {
		// sort each string first.
		str := []byte(s)
		slices.Sort(str)
		ss := string(str)

		// fmt.Println(ss)

		anagrams[ss] = append(anagrams[ss], s)
	}

	// fmt.Println(anagrams)

	ret := make([][]string, 0, len(anagrams))
	for _, v := range anagrams {
		ret = append(ret, v)
	}

	return ret
}
