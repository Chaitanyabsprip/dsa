package main

import "fmt"

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the
median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).



Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.

Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.



Constraints:

    nums1.length == m
    nums2.length == n
    0 <= m <= 1000
    0 <= n <= 1000
    1 <= m + n <= 2000
    -106 <= nums1[i], nums2[i] <= 106

*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var former, latter []int
	if len(nums1) >= len(nums2) {
		former = nums1
		latter = nums2
	}
	if len(nums1) < len(nums2) {
		former = nums2
		latter = nums1
	}
	pointer, left, right := 0, 0, len(former)

	/*

	   choose an element from latter
	   check if the element is less than current median
	   assume the element to be added and calculate new median

	   for all indexes above the index

	*/

	for left < right {
		target := latter[pointer]
		mid := left + (right-left)/2

		if former[mid] > target && former[mid-1] < target {
		}

		if former[mid] > latter[pointer] {
			right = mid
		} else if former[mid] < latter[pointer] {
			left = mid + 1
		}
	}

	return 0
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{}, []int{}))
}
