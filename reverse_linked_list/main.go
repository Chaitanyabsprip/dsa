package main

import "fmt"

/*
Given the head of a singly linked list, reverse the list, and return the
reversed list.

Example 1:

Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:

Input: head = [1,2]
Output: [2,1]

Example 3:

Input: head = []
Output: []



Constraints:

    The number of nodes in the list is the range [0, 5000].
    -5000 <= Node.val <= 5000
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func arrToLinkedList(arr []int) *ListNode {
	current := &ListNode{}
	head := struct{ head *ListNode }{current}
    head := new(ListNode)
    *head = *current

	for i, val := range arr {
		current.Val = val
		if i < len(arr)-1 {
			current.Next = &ListNode{}
		}
		current = current.Next
	}
	return head.head
}

func linkedListToArr(head *ListNode) []int {
	result := make([]int, 0)
	for ; head != nil; head = head.Next {
		result = append(result, head.Val)
	}
	return result
}

func main() {
	fmt.Println(linkedListToArr(reverseList(arrToLinkedList([]int{1, 2}))))
	fmt.Println(linkedListToArr(reverseList(arrToLinkedList([]int{1, 2, 3, 4, 5}))))
}
