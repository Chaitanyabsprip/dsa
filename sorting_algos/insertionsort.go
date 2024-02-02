package main

func insertionsort(array []int) []int {
	n := len(array)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			} else {
				break
			}
			j = j - 1
		}
	}
	return array
}
