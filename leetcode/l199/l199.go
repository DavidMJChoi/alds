// LeetCode L199

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	rsv := []int{}

	for len(queue) > 0 {
		rsv = append(rsv, queue[len(queue)-1].Val)
		newQueue := []*TreeNode{}
		for len(queue) > 0 {
			p := queue[0]
			if p != nil {
				if p.Left != nil {
					newQueue = append(newQueue, p.Left)
				}
				if p.Right != nil {
					newQueue = append(newQueue, p.Right)
				}
			}
			queue = queue[1:]
		}
		queue = newQueue
	}

	return rsv
}
