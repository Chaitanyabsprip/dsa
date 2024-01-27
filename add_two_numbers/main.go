package main

import (
	"math"
)

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a
single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the
number 0 itself.



Example 1:

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]



Constraints:

    The number of nodes in each linked list is in the range [1, 100].
    0 <= Node.val <= 9
	It is guaranteed that the list represents a number that does not have
	leading zeros.

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sumll := new(ListNode)
	save := sumll
	carry := 0
	var sum int
	for l1 != nil || l2 != nil {
		if l1 == nil {
			sum = l2.Val + carry
			l2 = l2.Next
		} else if l2 == nil {
			sum = l1.Val + carry
			l1 = l1.Next
		} else {
			sum = l1.Val + l2.Val + carry
			l1 = l1.Next
			l2 = l2.Next
		}
		if sum > 9 {
			sum = sum % 10
			carry = sum / 10
		}
		sumll.Val = sum
		sumll.Next = new(ListNode)
		sumll = sumll.Next
	}
	return save
}

func genNum(l *ListNode) int {
	var recurse func(*ListNode, int) int
	recurse = func(ln *ListNode, i int) int {
		if ln == nil {
			return 0
		}
		return ln.Val*int(math.Pow(10, float64(i))) + recurse(ln.Next, i+1)
	}
	return recurse(l, 1)
}

func main() {}
