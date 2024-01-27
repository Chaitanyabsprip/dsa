package main

import "fmt"

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	set := make(map[rune]int, len(s))
	for _, char := range s {
		set[char] += 1
	}
	for _, char := range t {
		set[char] -= 1
	}
	for _, count := range set {
		if count != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagalam"))
}
