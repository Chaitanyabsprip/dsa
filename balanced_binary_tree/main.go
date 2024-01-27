package main

/*
Given a binary tree, determine if it is height-balanced.

A height-balanced binary tree is a binary tree in which the depth of the two
subtrees of every node never differs by more than one.


Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: true

Example 2:

Input: root = [1,2,2,3,3,null,null,4,4]
Output: false

Example 3:

Input: root = []
Output: true



Constraints:

    The number of nodes in the tree is in the range [0, 5000].
    -104 <= Node.val <= 104

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	diff := abs(depth(root.Left) - depth(root.Right))
	return diff < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(depth(node.Left), depth(node.Right))
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
}
