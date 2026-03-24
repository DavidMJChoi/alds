// LeetCode L102

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	levels := [][]int{}

	for len(queue) > 0 {
		level := []int{}
		newQueue := []*TreeNode{}
		for len(queue) > 0 {
			p := queue[0]
			if p != nil {
				level = append(level, p.Val)
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
		if len(level) != 0 {
			levels = append(levels, level)
		}
	}

	return levels
}
