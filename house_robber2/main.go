package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	// dp[i] represents optimal value for nums[:i]
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i, n := range nums[2:] {
		dp[i+2] = max(
			dp[i]+n,
			dp[i+1],
		)
	}
	return dp[len(nums)-1]
}

func main() {
	fmt.Println(rob())
}
