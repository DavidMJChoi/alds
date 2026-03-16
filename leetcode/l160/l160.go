// LeetCode L160
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// solution 1: hash table
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	set := make(map[*ListNode]struct{})

	for headA != nil {
		set[headA] = struct{}{}
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := set[headB]; ok {
			return headB
		}
		headB = headB.Next
	}

	return nil
}

// solution 2: two alternativing pointers
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA := headA
	pB := headB

	for !(pA == nil && pB == nil) {
		if pA == pB {
			return pA
		}
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	return nil
}

// solution 3: compensate length difference
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// find the lengths first
	la, lb := 0, 0
	for p := headA; p != nil; p = p.Next {
		la++
	}
	for p := headB; p != nil; p = p.Next {
		lb++
	}

	// find the difference and compensate
	d := la - lb
	if d > 0 { // la is larger, compensate listA for letting it walk d steps first
		for d > 0 {
			headA = headA.Next
			d--
		}
	} else {
		for d < 0 {
			headB = headB.Next
			d++
		}
	}

	// now two linked list are synchronized
	for !(headA == nil && headB == nil) {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}
