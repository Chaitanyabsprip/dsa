package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sp := maxDepth(root.Left) + maxDepth(root.Right)
	return max(sp, diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func main() {
}
