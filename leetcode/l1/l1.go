// Keys: sort first. keep original indices, extra space.
// Time: O(nlogn) ordinary sorts
// Time: O(n) linear-time sorting for integers

package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 4}
	target := 6

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	hm := make(map[int]int)

	for i := range nums {
		n := nums[i]
		t := target - n
		j, ok := hm[t]
		if ok {
			return []int{i, j}
		} else {
			hm[n] = i
		}
	}

	return []int{}
}

func twoSumSort(nums []int, target int) []int {
	indices := radixSortIndices(nums)
	nums = getSortedByIndices(nums, indices)

	ret := []int{}

	i, j := 0, len(nums)-1
	for i <= j {
		t := target - nums[i]
		if nums[j] == t {
			ret = append(ret, indices[i])
			ret = append(ret, indices[j])
			break
		} else if nums[j] > t {
			j--
		} else {
			i++
		}
	}

	return ret
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
