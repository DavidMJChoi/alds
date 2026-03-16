// LeetCode L142

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head.Next

	for fast != nil {
		if slow == fast {
			slow = slow.Next
			p := head
			for p != slow {
				slow = slow.Next
				p = p.Next
			}
			return p
		}
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		slow = slow.Next
	}
	return nil
}
