package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]int, len(nums))
	for _, num := range nums {
		if seen[num] > 0 {
			return true
		}
		seen[num] += 1
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3}))
}
