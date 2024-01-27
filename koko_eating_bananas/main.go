package main

import (
	"fmt"
	"math"
	"slices"
)

func minEatingSpeed(piles []int, h int) int {
	mx := slices.Max(piles)
	var sum int
	for _, pile := range piles {
		sum += pile
	}
	mn := sum / h
	if mn == 0 {
		mn = 1
	}

	for mn < mx {
		mid := (mx + mn) / 2
		var temp float64
		for _, pile := range piles {
			temp += math.Ceil(float64(pile) / float64(mid))
		}
		if int(temp) <= h {
			mx = mid
		} else {
			mn = mid + 1
		}

	}
	return mn
}

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(minEatingSpeed([]int{312884470}, 968709470))
}
