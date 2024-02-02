package main

import (
	"fmt"
	"slices"
)

/*
In this problem, a tree is an undirected graph that is connected and has no
cycles.

You are given a graph that started as a tree with n nodes labeled from 1 to n,
with one additional edge added. The added edge has two different vertices chosen
from 1 to n, and was not an edge that already existed. The graph is represented
as an array edges of length n where edges[i] = [ai, bi] indicates that there is
an edge between nodes ai and bi in the graph.

Return an edge that can be removed so that the resulting graph is a tree of n
nodes. If there are multiple answers, return the answer that occurs last in the
input.



Example 1:

Input: edges = [[1,2],[1,3],[2,3]]
Output: [2,3]

Example 2:

Input: edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]
Output: [1,4]



Constraints:

    n == edges.length
    3 <= n <= 1000
    edges[i].length == 2
    1 <= ai < bi <= edges.length
    ai != bi
    There are no repeated edges.
    The given graph is connected.

*/

func findRedundantConnection(edges [][]int) []int {
	adjacent := make(map[int][]int)
	for _, edge := range edges {
		n1 := edge[0]
		n2 := edge[1]
		adjacent[n1] = append(adjacent[n1], n2)
		adjacent[n2] = append(adjacent[n2], n1)
	}

	fmt.Println(adjacent)

	cycle := dfs(edges[0][0], -1, adjacent, make([]int, 0), make(map[int]bool))
	i := slices.Index(cycle, cycle[len(cycle)-1])
	cycle = cycle[i:]

	fmt.Println(cycle)
	for i := len(edges) - 1; i > -1; i-- {
		edge := edges[i]
		if slices.Contains(cycle, edge[0]) && slices.Contains(cycle, edge[1]) {
			return edge
		}
	}
	fmt.Println("unreachable")

	return []int{}
}

func dfs(node, parent int, adjacent map[int][]int, path []int, seen map[int]bool) []int {
	fmt.Println("to", node, path, seen)
	if seen[node] {
		return append(path, node)
	}
	seen[node] = true
	path = append(path, node)
	for _, n := range adjacent[node] {
		if n != parent {
			nPath := dfs(n, node, adjacent, path, seen)
			if len(nPath) > len(path) {
				return nPath
			}
		}
	}
	path = path[:len(path)-1]
	return path
}

func main() {
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}))
}
