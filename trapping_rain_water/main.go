package main

import (
	"fmt"
)

/*
Given n non-negative integers representing an elevation map where the width of
each bar is 1, compute how much water it can trap after raining.



Example 1:

Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array
[0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section)
are being trapped.

Example 2:

Input: height = [4,2,0,3,2,5]
Output: 9



Constraints:

    n == height.length
    1 <= n <= 2 * 104
    0 <= height[i] <= 105

*/

func maxima(slice []int) int {
	mx := slice[0]
	idx := 0
	for i, elem := range slice {
		if elem > mx {
			mx = elem
			idx = i
		}
	}
	return idx
}

func trap(height []int) int {
	left, right := 0, 1
	var area int
	mxi := maxima(height)
	for right < mxi {
		if height[right] < height[left] {
			fmt.Println(left, right, height[left]-height[right])
			area += height[left] - height[right]
		} else {
			left = right
		}
		right++
	}

	left, right = len(height)-1, len(height)-2
	for right > mxi {
		if height[right] < height[left] {
			fmt.Println(left, right, height[left]-height[right])
			area += height[left] - height[right]
		} else {
			left = right
		}
		right--
	}
	return area
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}
