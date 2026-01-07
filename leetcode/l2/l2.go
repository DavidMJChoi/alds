// Keys: simply converts to ints/doubles and adding up will probably fail since the numbers are too large (10^100)

package main

import "fmt"

func main() {

	l1 := sliceToList([]int{2, 4, 3}) // 342
	l2 := sliceToList([]int{5, 6, 4}) // 465

	sumList := addTwoNumbers(l1, l2)

	fmt.Println(listToSlice(sumList))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := listToSlice(l1)
	s2 := listToSlice(l2)

	return sliceToList(sliceSum(s1, s2))
}

func listToSlice(l *ListNode) []int {
	ret := []int{}
	for l != nil {
		ret = append(ret, l.Val)
		l = l.Next
	}
	return ret
}

func sliceToList(num []int) *ListNode {
	l := &ListNode{num[0], nil}
	root := l

	for i := 1; i < len(num); i++ {
		l.Next = &ListNode{num[i], nil}
		l = l.Next
	}

	return root
}

func sliceSum(s1, s2 []int) []int {
	if len(s1) < len(s2) {
		s1 = append(s1, make([]int, len(s2)-len(s1))...)
	} else {
		s2 = append(s2, make([]int, len(s1)-len(s2))...)
	}

	ret := make([]int, len(s1))

	carry := 0
	for i := 0; i < len(s1); i++ {
		sum := s1[i] + s2[i] + carry
		ret[i] = sum % 10
		carry = sum / 10
	}

	if carry != 0 {
		ret = append(ret, carry)
	}

	return ret
}
