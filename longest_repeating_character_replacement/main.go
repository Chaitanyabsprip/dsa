package main

import "fmt"

func characterReplacement(s string, k int) int {
	mx := 0

	for width := k + 1; width < len(s); width-- {
	}
	for j := range s {
		tmx := 0
		kc := k
		for i := j; i < len(s); i++ {
			if s[i] != s[j] {
				if kc == 0 {
					break
				}
				kc--
			}
			tmx++
		}
		mx = max(mx, tmx)
	}
	return mx
}

func main() {
	fmt.Println(characterReplacement("ABAB", 2))    // 4
	fmt.Println(characterReplacement("ABBABBA", 2)) // 4
}
