// LeetCode L23

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		newLists := []*ListNode{}
		for i := 0; i < len(lists); i += 2 {
			if i+1 < len(lists) {
				newLists = append(newLists, mergeTwoLists(lists[i], lists[i+1]))
			} else {
				newLists = append(newLists, lists[i])
			}
		}
		lists = newLists
	}
	return lists[0]
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := new(ListNode)
	prev := preHead
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	if list1 == nil {
		prev.Next = list2
	} else {
		prev.Next = list1
	}
	return preHead.Next
}
