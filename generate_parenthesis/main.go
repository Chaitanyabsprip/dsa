package main

import "fmt"

/*
Given n pairs of parentheses, write a function to generate all combinations of
well-formed parentheses.



Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:

Input: n = 1
Output: ["()"]



Constraints:

    1 <= n <= 8

*/

func generateParenthesis(n int) []string {
	stack := make([]string, 0)
	res := make([]string, 0)
	recurse(n, "", stack, &res)
	return res
}

func recurse(n int, curr string, stack []string, res *[]string) {
	if len(curr) == 2*n && len(stack) == 0 {
		*res = append(*res, curr)
		return
	}
	if len(stack) < n && len(curr) < 2*n-1 {
		recurse(n, curr+"(", append(stack, ")"), res)
	}
	if len(stack) > 0 {
		recurse(n, curr+")", stack[:len(stack)-1], res)
	}
}

func main() {
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
}
