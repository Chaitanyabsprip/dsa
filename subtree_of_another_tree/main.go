package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return true
    }
    if root.Val == subRoot.Val {
        return isSameTree(root, subRoot)
    }
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil {
        return q == nil
    }
    if q == nil {
        return p == nil
    }
    if p.Val != q.Val {
        return false
    }
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	fmt.Println(isSubtree())
}
