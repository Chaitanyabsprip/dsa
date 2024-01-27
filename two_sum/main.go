package main

import "fmt"

func twoSum(nums []int, target int) []int {
	set := make(map[int]int)
	for i, num := range nums {
		if j, hasNum := set[target - num]; hasNum {
			return []int{i, j}
		}
		set[num] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6)
}
