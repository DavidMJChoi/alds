// LeetCode L138

package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	// treat the evil [] as a edge case
	if head == nil {
		return nil
	}

	// get list length
	l := 0
	for p := head; p != nil; p = p.Next {
		l++
	}

	// construct jumpList
	jumpList := make([]int, l)
	p := head
	i := 0
	for p != nil {
		if p.Random == nil {
			jumpList[i] = -1
		} else {
			q := p.Random
			m := 0
			for q != nil {
				m++
				q = q.Next
			}
			jumpList[i] = l - m
		}
		p = p.Next
		i++
	}

	// deep copy, ignoring p.Random
	h := new(Node)
	p = h
	q := head
	for q != nil {
		p.Val = q.Val
		if q.Next != nil {
			p.Next = new(Node)
		}
		p = p.Next
		q = q.Next
	}

	// use jumpList to construct p.Random
	i = 0
	p = h
	for p != nil {
		if jumpList[i] == -1 {
			p.Random = nil
		} else {
			q := h
			for range jumpList[i] {
				q = q.Next
			}
			p.Random = q
		}
		p = p.Next
		i++
	}

	return h
}
