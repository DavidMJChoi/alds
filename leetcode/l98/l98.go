// LeetCode L98

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isValidBST(root *TreeNode) bool {
	// in-order traversal -> ordered list

	visited := make(map[*TreeNode]struct{})
	stack := []*TreeNode{root}

	// find smallest value first
	p := root
	for p.Left != nil {
		p = p.Left
	}
	prevValue := p.Val
	skippedFirst := false

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		_, ok := visited[p]
		if p.Left == nil && p.Right == nil || ok {
			if skippedFirst && p.Val <= prevValue {
				return false
			} else {
				prevValue = p.Val
			}
			if !skippedFirst {
				skippedFirst = true
			}
		} else {
			visited[p] = struct{}{}
			if p.Right != nil {
				stack = append(stack, p.Right)
			}
			stack = append(stack, p)
			if p.Left != nil {
				stack = append(stack, p.Left)
			}
		}
	}

	return true
}
