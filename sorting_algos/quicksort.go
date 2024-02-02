package main

import "fmt"

// how to handle repeated elements without the equal array?
func quicksort(array []int) []int {
	if len(array) == 0 {
		return array
	}
	pivot := len(array) / 2
	fmt.Println(pivot, array)
	left := make([]int, 0)
	right, equal := left, left
	for _, n := range array {
		if n < array[pivot] {
			left = append(left, n)
		} else if n > array[pivot] {
			right = append(right, n)
		} else if n == array[pivot] {
			equal = append(equal, n)
		}
	}

	return append(append(quicksort(left), equal...), quicksort(right)...)
}
