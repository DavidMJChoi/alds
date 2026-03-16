// LeetCode L141

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// solution1: hash set
func hasCycle1(head *ListNode) bool {
	set := make(map[*ListNode]struct{})

	for head != nil {
		if _, ok := set[head]; ok {
			return true
		} else {
			set[head] = struct{}{}
		}
		head = head.Next
	}

	return false
}

// solution2: fast and slow pointers
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}

	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		}

	}
	return false
}
