// LeetCode L148

package main

import "slices"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// get list length
	l := 0
	for p := head; p != nil; p = p.Next {
		l++
	}

	// list to slice
	s := make([]int, l)
	i := 0
	for p := head; p != nil; p = p.Next {
		s[i] = p.Val
		i++
	}

	slices.Sort(s)

	// slice to list

	h := new(ListNode)
	p := h
	i = 0
	for i < l {
		p.Val = s[i]
		if i != l-1 {
			p.Next = new(ListNode)
			p = p.Next
		}
		i++
	}

	return h
}
