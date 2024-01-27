package main

import "fmt"

/*
Given an array of integers nums which is sorted in ascending order, and an
integer target, write a function to search target in nums. If target exists,
then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.



Example 1:

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Example 2:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1



Constraints:

    1 <= nums.length <= 104
    -104 < nums[i], target < 104
    All the integers in nums are unique.
    nums is sorted in ascending order.

*/

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (right + left) / 2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{1}, 1))                   // 0
	fmt.Println(search([]int{1, 3}, 1))                // 0
	fmt.Println(search([]int{1, 3}, 3))                // 1
	fmt.Println(search([]int{1, 3, 4, 5, 2, 6, 9}, 5)) // 3
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))  // 4
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))  // -1
}
