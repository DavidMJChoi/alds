// LeetCode L230

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func kthSmallest(root *TreeNode, k int) int {
	// pre-order traversal -> ordered list
	visited := make(map[*TreeNode]struct{})
	stack := []*TreeNode{root}

	i := 1
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		_, ok := visited[p]
		if p.Left == nil && p.Right == nil || ok {
			if i == k {
				return p.Val
			}
			i++
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

	return 0
}
