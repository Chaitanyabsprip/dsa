package main

import (
	"math"
	"slices"
)

/*
Design a stack that supports push, pop, top, and retrieving the minimum element
in constant time.

Implement the MinStack class:

    MinStack() initializes the stack object.
    void push(int val) pushes the element val onto the stack.
    void pop() removes the element on the top of the stack.
    int top() gets the top element of the stack.
    int getMin() retrieves the minimum element in the stack.

You must implement a solution with O(1) time complexity for each function.



Example 1:

Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2



Constraints:

    -231 <= val <= 231 - 1
	Methods pop, top and getMin operations will always be called on non-empty
	stacks.
    At most 3 * 104 calls will be made to push, pop, top, and getMin.

*/

type MinStack struct {
	stack []int
	min   int
}

func Constructor() MinStack {
	s := new(MinStack)
	s.stack = make([]int, 0)
	s.min = math.MaxInt32
	return *s
}

func (this *MinStack) Push(val int) {
	this.min = min(this.min, val)
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.min = slices.Min(this.stack)
	if len(this.stack) == 0 {
		this.min = math.MaxInt32
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
}
