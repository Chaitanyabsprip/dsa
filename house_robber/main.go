package main

import (
	"fmt"
	"slices"
)

func rob(nums []int) int {
	memo := make(map[int]int)
	return fun(nums, memo)
}

func fun(nums []int, memo map[int]int) int {
	if val, ok := memo[len(nums)]; ok {
		return val
	}
	var result int
	if len(nums) <= 2 {
		result = slices.Max(nums)
	} else if len(nums) == 3 {
		result = max(nums[0]+nums[2], nums[1])
	} else {
		zeroIndexSum := nums[0] + fun(nums[2:], memo)
		firstIndexSum := nums[1] + rob(nums[3:])
		result = max(zeroIndexSum, firstIndexSum)
	}
	memo[len(nums)] = result
	return result
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1})) // 4
}
