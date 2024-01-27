package main

/*
Given an m x n matrix board containing 'X' and 'O', capture all regions that are
4-directionally surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.



Example 1:

Input: board =
[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
Output: [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
Explanation: Notice that an 'O' should not be flipped if:
- It is on the border, or
- It is adjacent to an 'O' that should not be flipped.
The bottom 'O' is on the border, so it is not flipped.
The other three 'O' form a surrounded region, so they are flipped.

Example 2:

Input: board = [["X"]]
Output: [["X"]]



Constraints:

    m == board.length
    n == board[i].length
    1 <= m, n <= 200
    board[i][j] is 'X' or 'O'.

*/

func solve(board [][]byte) {
	seen := new([][]bool)
	n := len(board)
	m := len(board[0])
	*seen = make([][]bool, n)
	for i := 0; i < n; i++ {
		(*seen)[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 || i == n-1 || j == m-1 {
				recurse(board, i, j, seen)
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'T' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func recurse(board [][]byte, i, j int, seen *[][]bool) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || (board)[i][j] != 'O' {
		return
	}
	board[i][j] = 'T'
	recurse(board, i, j-1, seen)
	recurse(board, i, j+1, seen)
	recurse(board, i-1, j, seen)
	recurse(board, i+1, j, seen)
}

func main() {
}
