// LeetCode L206

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

type stk struct {
	sl []*ListNode
}

func (s *stk) push(n *ListNode) {
	s.sl = append(s.sl, n)
}

func (s *stk) pop() *ListNode {
	if s.isEmpty() {
		return nil
	}
	ret := s.sl[len(s.sl)-1]
	s.sl = s.sl[:len(s.sl)-1]
	return ret
}

func (s *stk) Len() int {
	return len(s.sl)
}

func (s *stk) isEmpty() bool {
	return len(s.sl) == 0
}

func reverseList1(head *ListNode) *ListNode {
	s := new(stk)

	if head == nil {
		return head
	}

	p := head
	for p.Next != nil {
		s.push(p)
		t := p
		p = p.Next
		t.Next = nil
	}

	head = p
	for !s.isEmpty() {
		p.Next = s.pop()
		p = p.Next
	}

	return head
}

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	prev := head
	curr := prev.Next
	prev.Next = nil
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
