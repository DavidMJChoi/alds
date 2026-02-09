// LeetCode L28
// Of course this is easy... If you can tolerate O(n^2) time complexity
// To achive a O(m+n) complexity, use KMP or a similar algorithm.
// I find KMP tricky to understand. I am very stupid so you might find different since you are probably smarter.

package main

import "fmt"

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
	fmt.Println(strStr("sadbutsad", "leeto"))
}

func strStr(haystack string, needle string) int {
	return KMP(haystack, needle)
}

func KMP(T, P string) int {
	n, m := len(T), len(P)
	if m == 0 {
		return 0 // 空模式串匹配任何字符串的起始位置
	}

	pmt := buildPMT(P) // 构建部分匹配表（失败函数）
	i, j := 0, 0       // i 是 T 的索引，j 是 P 的索引

	for i < n {
		// 如果字符匹配，继续比较下一个字符
		if T[i] == P[j] {
			i++
			j++
			// 如果完全匹配，返回匹配的起始位置
			if j == m {
				return i - m
			}
		} else {
			// 如果字符不匹配，根据部分匹配表回退 j
			if j > 0 {
				j = pmt[j-1]
			} else { // failed at the first character of P
				i++
			}
		}
	}

	return -1 // 未找到匹配
}

func buildPMT(P string) []int {
	m := len(P)
	pmt := make([]int, m)
	pmt[0] = 0 // 第一个字符的前缀和后缀长度为 0
	i := 1
	j := 0

	for i < m {
		if P[i] == P[j] {
			pmt[i] = j + 1
			i++
			j++
		} else if j == 0 {
			pmt[i] = 0
			i++
		} else {
			j = pmt[j-1]
		}
	}

	return pmt
}
