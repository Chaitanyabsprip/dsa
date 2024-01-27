package main

import "fmt"

/*
You are given an m x n integer matrix matrix with the following two properties:

    Each row is sorted in non-decreasing order.
	The first integer of each row is greater than the last integer of the
	previous row.

Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.



Example 1:

Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true

Example 2:

Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false



Constraints:

    m == matrix.length
    n == matrix[i].length
    1 <= m, n <= 100
    -104 <= matrix[i][j], target <= 104

*/

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])
	var lower, upper, mid int
	lower = 0
	upper = (n - 1) * (m - 1)
	if upper == 0 {
		return matrix[0][0] == target
	}
	for lower < upper {
		mid = lower + (upper-lower)/2
		i := mid / n
		j := mid % n
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			lower = mid + 1
		} else if matrix[i][j] > target {
			upper = mid - 1
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))  // true
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13)) // false
}
