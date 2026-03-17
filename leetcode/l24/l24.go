// LeetCode L24

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func swapPairs(head *ListNode) *ListNode {
	p := head
	// modify head if necessary
	if p != nil && p.Next != nil {
		head = p.Next
	}

	var prevP *ListNode = nil
	for p != nil {
		if q := p.Next; q != nil {
			next := q.Next
			q.Next = p
			p.Next = next
			if prevP != nil {
				prevP.Next = q
			}
			prevP = p
		}
		p = p.Next
	}
	return head
}
