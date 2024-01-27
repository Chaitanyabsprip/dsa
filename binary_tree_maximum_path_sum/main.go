package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := 0
	pathSum := func(node *TreeNode) int {
		if root == nil {
			return 0
		}
		sum := node.Val + max(pathSum(node.Left), pathSum(node.Right))
		return sum
	}
	return pathSum(root)
	return max(pathSum(root.Left)+pathSum(root.Right), maxPathSum(root.Left), maxPathSum(root.Right))
}

func pathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + max(pathSum(root.Left), pathSum(root.Right))
}

func main() {
	fmt.Println(maxPathSum())
}
