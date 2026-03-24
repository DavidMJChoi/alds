// LeetCode L114

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func flatten(root *TreeNode) {
	p := root
	for p != nil {
		r := p.Right
		p.Right = p.Left
		p.Left = nil

		// find the grafting node for r
		q := p
		for q.Right != nil {
			q = q.Right
		}
		q.Right = r

		p = p.Right
	}
}
