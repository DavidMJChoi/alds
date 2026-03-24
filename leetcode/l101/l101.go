// LeetCode L101

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isSymmetric(root *TreeNode) bool {
	return sym(root.Left, root.Right)
}

func sym(l, r *TreeNode) bool {

	if l == nil && r == nil {
		return true
	}

	if l == nil && r != nil || l != nil && r == nil {
		return false
	}

	if l.Val != r.Val {
		return false
	}

	if l.Left == nil && l.Right == nil && r.Left == nil && r.Right == nil {
		return l.Val == r.Val
	}

	return sym(l.Left, r.Right) && sym(l.Right, r.Left)
}
