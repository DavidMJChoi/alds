// LeetCode L25

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reverseKGroup(head *ListNode, k int) *ListNode {
	l := 0
	for r := head; r != nil; r = r.Next {
		l++
	}
	// p is the head of k-group
	// q is the tail of k-group
	// prevP is the tail of the last k-group
	p := head
	q := head
	var prevP *ListNode

	// find new head. No need to worry about nil since n >= k
	for range k - 1 {
		head = head.Next
	}

	// we only need to perfrom reverse k-group for l/k times
	for range l / k {
		i := 1
		var curr *ListNode
		for i < k {
			curr = q.Next
			next := curr.Next
			curr.Next = p
			p = curr
			q.Next = next
			i++
		}
		// til this point, p -> q is a reversed k-group
		// prevP is nil for the first k-group
		if prevP != nil {
			prevP.Next = p
		}
		prevP = q
		q = q.Next
		p = q
	}

	return head
}
