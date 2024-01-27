package main

/*
Given the root of a binary tree, imagine yourself standing on the right side of
it, return the values of the nodes you can see ordered from top to bottom.

Example 1:

Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]

Example 2:

Input: root = [1,null,3]
Output: [1,3]

Example 3:

Input: root = []
Output: []

Constraints:

    The number of nodes in the tree is in the range [0, 100].
    -100 <= Node.val <= 100
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	arr := make([][]int, 0)
	var fun func(*TreeNode, int)
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
	result := make([]int, 0)
	for _, larr := range arr {
		result = append(result, larr[len(larr)-1])
	}
	return result
}

func main() {}
