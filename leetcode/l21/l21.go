// LeetCode L21

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	p := list1
	q := list2
	var mergedList *ListNode = new(ListNode)
	r := mergedList
	for p != nil && q != nil {
		if p.Val > q.Val {
			r.Val = q.Val
			q = q.Next
		} else {
			r.Val = p.Val
			p = p.Next
		}

		r.Next = new(ListNode)
		r = r.Next
	}

	for p != nil {
		r.Val = p.Val
		p = p.Next
		if p != nil {
			r.Next = new(ListNode)
		}
		r = r.Next
	}
	for q != nil {
		r.Val = q.Val
		q = q.Next
		if q != nil {
			r.Next = new(ListNode)
		}
		r = r.Next
	}

	return mergedList
}
