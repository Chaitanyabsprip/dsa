package main

import "fmt"

func key(str string) [26]byte {
	hash := [26]byte{}
	for _, k := range str {
		hash[k-'a']++
	}
	return hash
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]byte][]string)
	for _, str := range strs {
		hash := key(str)
		groups[hash] = append(groups[hash], str)
	}
	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}
	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
