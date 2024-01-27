package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	seen := make(map[string]int)
	for i := 0; i < len(s)-9; i++ {
		seen[s[i:i+10]]++
	}
	result := make([]string, 0)
	for k, v := range seen {
		if v > 1 {
			result = append(result, k)
		}
	}
	return result
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")) // ["AAAAACCCCC","CCCCCAAAAA"]
}
