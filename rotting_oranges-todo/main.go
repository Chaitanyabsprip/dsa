package main

/*
You are given an m x n grid where each cell can have one of three values:

0 representing an empty cell, 1 representing a fresh orange, or 2 representing a
rotten orange.

Every minute, any fresh orange that is 4-directionally adjacent to a rotten
orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh
orange. If this is impossible, return -1.



Example 1:

Input: grid = [[2,1,1],[1,1,0],[0,1,1]] Output: 4

Example 2:

Input: grid = [[2,1,1],[0,1,1],[1,0,1]] Output: -1 Explanation: The orange in
the bottom left corner (row 2, column 0) is never rotten, because rotting only
happens 4-directionally.

Example 3:

Input: grid = [[0,2]] Output: 0 Explanation: Since there are already no fresh
oranges at minute 0, the answer is just 0.



Constraints:

m == grid.length n == grid[i].length 1 <= m, n <= 10 grid[i][j] is 0, 1, or 2.

*/

func orangesRotting(grid [][]int) int {
	distance := make([][]int, len(grid))
	seen := make([][]bool, len(grid))
	for i := range distance {
		distance[i] = make([]int, len(grid[i]))
		seen[i] = make([]bool, len(grid[i]))
		for j := range distance[i] {
			distance[i][j] = 200
		}
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				seenReset(&seen)
				grid[i][j] = 1
				fun(grid, i, j, 0, &distance, seen)
			}
		}
	}
	mx := 0
	for i := range distance {
		for j := range distance[i] {
			if grid[i][j] == 1 {
				mx = max(mx, distance[i][j])
			}
		}
	}
	if mx == 200 {
		return -1
	}
	return mx
}

func seenReset(seen *[][]bool) {
	for i := range *seen {
		for j := range (*seen)[i] {
			(*seen)[i][j] = false
		}
	}
}

func fun(grid [][]int, i, j, prev int, distance *[][]int, seen [][]bool) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 || seen[i][j] {
		return
	}
	(*distance)[i][j] = min(prev, (*distance)[i][j])
	seen[i][j] = true
	go fun(grid, i, j-1, prev+1, distance, seen)
	go fun(grid, i, j+1, prev+1, distance, seen)
	go fun(grid, i-1, j, prev+1, distance, seen)
	go fun(grid, i+1, j, prev+1, distance, seen)
}

func bfs(grid [][]int) int {
	queue := make([][]int, 0)
	ptr := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j, 0})
			}
		}
	}

	for ptr != len(queue) {
		q := queue[ptr]
		i, j, distance := q[0], q[1], q[2]
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
			ptr++
			continue
		}
		if grid[i][j] == 1 {
			grid[i][j] = 2
			queue = append(queue, []int{i - 1, j, distance + 1})
			queue = append(queue, []int{i + 1, j, distance + 1})
			queue = append(queue, []int{i, j - 1, distance + 1})
			queue = append(queue, []int{i, j + 1, distance + 1})
		}
		ptr++
	}
	v := queue[len(queue)-1]
	return v[len(v)-1]
}

/*
2 1 1
1 1 1
0 1 2

2 1 1
1 1 1
0 1 2
*/

func main() {}
