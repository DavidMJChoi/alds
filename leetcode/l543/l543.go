// LeetCode L543

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func diameterOfBinaryTree(root *TreeNode) int {

	ans := 1

	d, ans := depth(root, ans)
	d += 1

	return ans - 1
}

func depth(node *TreeNode, ans int) (int, int) {
	if node == nil {
		return 0, ans
	}

	L, ans1 := depth(node.Left, ans)
	R, ans2 := depth(node.Right, ans)
	ans = max(ans, ans1, ans2, L+R+1)
	return max(L, R) + 1, ans
}
