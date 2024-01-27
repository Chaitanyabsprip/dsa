package main

/*
Given the root of a binary tree, return the level order traversal of its nodes'
values. (i.e., from left to right, level by level).

Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Example 2:

Input: root = [1]
Output: [[1]]

Example 3:

Input: root = []
Output: []

Constraints:

    The number of nodes in the tree is in the range [0, 2000].
    -1000 <= Node.val <= 1000
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	arr := make([][]int, 0)
	var fun func(node *TreeNode, level int)
	fun = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(arr) <= level {
			arr = append(arr, make([]int, 0))
		}
		arr[level] = append(arr[level], node.Val)
		fun(node.Left, level+1)
		fun(node.Right, level+1)
	}
	fun(root, 0)
	return arr
}

func main() {
}
