// LeetCode L15
// Brute-force: sort, for each element, find two-sum. use a hash map to  remove duplicates. very slow...

package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// s1 := []int{-1, 0, 1, 2, -1, -4}
	// s2 := []int{0, 1, 1}
	// s3 := []int{0, 0, 0}
	s4 := []int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10}

	// fmt.Println(threeSum(s1))
	// fmt.Println(threeSum(s2))
	// fmt.Println(threeSum(s3))
	fmt.Println(threeSum(s4))

	// fmt.Println(sliceToString([]int{-2, 0, 2}))
}

func threeSum(nums []int) [][]int {
	nums = getSortedByIndices(nums, radixSortIndices(nums))

	ret := [][]int{}

	m := make(map[string]bool)

	// fmt.Println(nums)

	k := 0
	for k < len(nums) {
		i := 0

		// if k == 0 {
		// 	i++
		// }
		j := len(nums) - 1
		// if k == len(nums)-1 {
		// 	j--
		// }

		for i < j {
			if i == k {
				i++
				continue
			}
			if j == k {
				j--
				continue
			}

			// fmt.Println(i, j)

			sum3 := nums[i] + nums[j] + nums[k]

			// fmt.Println(nums[i], nums[k], nums[j])

			if sum3 == 0 {

				triplet := []int{nums[i], nums[j], nums[k]}
				slices.Sort(triplet)

				str := sliceToString(triplet)
				existed, _ := m[str]

				// fmt.Println(str)

				// fmt.Println(existed)

				if !existed {
					m[str] = true
					ret = append(ret, triplet)
				}

				i++
				j--
				continue
			} else {
				if sum3 > 0 {
					j--
				} else {
					i++
				}
			}
		}
		k++
	}

	return ret
}

// func sliceToString(s []int) string {
// 	ret := ""
// 	for _, v := range s {
// 		ret += strconv.Itoa(v)
// 	}

// 	return ret
// }

func sliceToString(s []int) string {
	if len(s) == 0 {
		return ""
	}

	var b strings.Builder
	b.Grow(len(s) * 7) // 假设每个数字平均占3字符，根据实际情况调整

	for _, v := range s {
		b.WriteString(strconv.Itoa(v))
	}

	return b.String()
}

func radixSortIndices(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}

	// 初始化索引数组
	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	// 将数值转为无符号整数（处理负数）
	uarr := make([]uint32, n)
	for i, v := range nums {
		// 翻转符号位，使负数在无符号比较时顺序正确
		uarr[i] = uint32(v) ^ 0x80000000
	}

	// 临时数组
	tempIndices := make([]int, n)
	tempValues := make([]uint32, n)

	// 按字节（0~3 字节）进行计数排序
	for shift := uint(0); shift < 32; shift += 8 {
		// 计数数组，每轮 256 个桶
		count := [256]int{}

		// 统计当前字节出现次数
		for _, v := range uarr {
			b := (v >> shift) & 0xFF
			count[b]++
		}

		// 前缀和，使计数排序稳定
		for i := 1; i < 256; i++ {
			count[i] += count[i-1]
		}

		// 从后往前遍历，放入临时数组（保持稳定性）
		for i := n - 1; i >= 0; i-- {
			b := (uarr[i] >> shift) & 0xFF
			count[b]--
			pos := count[b]
			tempValues[pos] = uarr[i]
			tempIndices[pos] = indices[i]
		}

		// 交换回主数组，准备下一轮
		uarr, tempValues = tempValues, uarr
		indices, tempIndices = tempIndices, indices
	}

	return indices
}

func getSortedByIndices(nums []int, indices []int) []int {
	sorted := make([]int, len(nums))
	for i, idx := range indices {
		sorted[i] = nums[idx]
	}
	return sorted
}
