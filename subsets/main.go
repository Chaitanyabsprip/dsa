package main

import "fmt"

/*
Given an integer array nums of unique elements, return all possible
subsets
(the power set).

The solution set must not contain duplicate subsets. Return the solution in any
order.



Example 1:

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:

Input: nums = [0]
Output: [[],[0]]



Constraints:

    1 <= nums.length <= 10
    -10 <= nums[i] <= 10
    All the numbers of nums are unique.

*/

func subsets(nums []int) [][]int {
	return backtrack(nums, []int{}, [][]int{{}}, 1)
	// return bt(nums, [][]int{{}})
}

func bt(nums []int, result [][]int) [][]int {
	if len(nums) == 0 {
		return result
	}
	for _, set := range result {
		result = append(result, append(set, nums[0]))
	}
	return bt(nums[1:], result)
}

func backtrack(nums, curr []int, result *[][]int, n int) {
	if len(nums) == 0 {
		*result = append(*result, curr)
	}
	for _, num := range nums {
		if n > 1 {
			backtrack(nums[1:], append(curr, nums[0]), result, n)
			backtrack(nums[1:], curr, result, n)
		}
	}
}

/*

[1, 2, 3]
---------

[]
[[] [1]]
[[] [1] [2] [1, 2]]
[[] [1] [2] [1, 2] [3] [1, 3] [2, 3] [1, 2, 3]]

[]
[[]]
[[] [1] [2] [3]]
[[] [1] [2] [3] [1, 2] [2, 3] [1, 3]]
[[] [1] [2] [3] [1, 2] [2, 3] [1, 3] [1, 2, 3]]


*/

// func backtrack(nums []int, result [][]int, curr []int, k int) {
// 	if len(curr) == k { // if goal is reached add to result
// 		result = append(result, curr)
// 		return
// 	}
//
// 	for i := 0; i < len(result); i++ {
// 		if choices[i] is valid {
// 			make choices[i]
// 			backtrack(num, result)
// 			undo choices[i]
// 		}
// 	}
// }

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
