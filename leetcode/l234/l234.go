// LeetCode L234

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func isPalindrome1(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	// find length
	l := 0
	p := head
	for p != nil {
		l++
		p = p.Next
	}
	// find the head for the second half
	mid := head
	if l%2 != 0 {
		mid = mid.Next
	}
	l /= 2
	for l > 0 {
		mid = mid.Next
		l--
	}

	// reverse second half
	prev := mid
	curr := prev.Next
	prev.Next = nil
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	head2 := prev

	// compare
	for head != nil && head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}

func isPalindrome(head *ListNode) bool {
	nums := []int{}

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	for i := range len(nums) / 2 {
		if nums[i] != nums[len(nums)-i-1] {
			return false
		}
	}
	return true
}
