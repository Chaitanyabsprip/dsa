package main

import (
	"fmt"
	"math"
)

/*
 */

func coinChange(coins []int, amount int) int {
	memo := make(map[int]int)
	val := fun(coins, amount, memo)
	// fmt.Println(memo)
	return val
}

func fun(coins []int, amount int, memo map[int]int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if val, ok := memo[amount]; ok {
		return val
	}
	mn := math.MaxInt32
	for _, coin := range coins {
		result := fun(coins, amount-coin, memo)
		if result != -1 {
			mn = min(mn, result+1)
		}
	}
	memo[amount] = mn
	if mn != math.MaxInt32 {
		return mn
	}
	return -1
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))             // 3
	fmt.Println(coinChange([]int{2}, 3))                    // -1
	fmt.Println(coinChange([]int{1}, 0))                    // 0
	fmt.Println(coinChange([]int{2}, 1))                    // -1
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249)) // 20
}

/*
11 with [1, 2, 5]

F(11) = min(F(10), F(9), F(6)) + 1
F(10) = min(F(9), F(8), F(5)) + 1
F(9) = min(F(8)+1, F(7)+1, F(4)+1)
F(8) = min(F(7)+1, F(6)+1, F(3)+1)
F(7) = min(F(6)+1, F(5)+1, F(2)+1)
F(6) = min(F(5)+1, F(4)+1, F(1)+1)
F(5) = min(F(4)+1, F(3)+1, F(0)+1)
F(4) = min(F(3)+1, F(2)+1, -1+1)
F(3) = min(F(2)+1, F(1)+1, -1+1)
F(2) = min(F(1)+1, F(0)+1, 1+1)
F(1) = min(F(0)+1, -1+1, -1+1)
F(0) = 0
F(-1) = -1



*/
