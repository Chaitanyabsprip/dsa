package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	hm := make(map[int]int, len(nums))
	for _, v := range nums {
		hm[v] += 1
	}
	// convert map to list of tuples of k, v
	ctr := make([][]int, 0)
	for k, v := range hm {
		ctr = append(ctr, []int{k, v})
	}
	// sort the slice of tuples of k, v based on v(freq)
	sort.Slice(ctr, func(i int, j int) bool {
		return ctr[i][1] > ctr[j][1]
	})

	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = ctr[i][0]
	}
	return ans
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
