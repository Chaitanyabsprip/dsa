package main

import "fmt"

/*
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := 0
	lenB := 0

	pa := headA
	for pa != nil {
		pa = pa.Next
		lenA++
	}
	pb := headB
	for pb != nil {
		pb = pb.Next
		lenB++
	}

	pa = headA
	pb = headB
	var short int
	if lenA > lenB {
		short = lenB
		for i := 0; i < lenA-lenB; i++ {
			pa = pa.Next
		}
	} else {
		short = lenA
		for i := 0; i < lenB-lenA; i++ {
			pb = pb.Next
		}
	}
	for i := 0; i < short; i++ {
		if pa == pb {
			return pa
		}
		pa = pa.Next
		pb = pb.Next
	}
	return nil
}

func main() {
}
