// LeetCode L19

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := head
	for range n {
		p = p.Next
	}
	q := head
	prevq := new(ListNode)

	for p != nil {
		p = p.Next
		prevq = q
		q = q.Next
	}

	// now q points to the element to be removed
	if q == head {
		return q.Next
	}
	prevq.Next = q.Next
	q.Next = nil

	return head
}
