package main

import "fmt"

/*
There are a total of numCourses courses you have to take, labeled from 0 to
numCourses - 1. You are given an array prerequisites where prerequisites[i] =
[ai, bi] indicates that you must take course bi first if you want to take course
ai.

	For example, the pair [0, 1], indicates that to take course 0 you have to
	first take course 1.

Return the ordering of courses you should take to finish all courses. If there
are many valid answers, return any of them. If it is impossible to finish all
courses, return an empty array.



Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should
have finished course 0. So the correct course order is [0,1].

Example 2:

Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should
have finished both courses 1 and 2. Both courses 1 and 2 should be taken after
you finished course 0.
So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].

Example 3:

Input: numCourses = 1, prerequisites = []
Output: [0]



Constraints:

    1 <= numCourses <= 2000
    0 <= prerequisites.length <= numCourses * (numCourses - 1)
    prerequisites[i].length == 2
    0 <= ai, bi < numCourses
    ai != bi
    All the pairs [ai, bi] are distinct.

*/

func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses == 0 {
		return []int{}
	}
	nodes := make(map[int][]int)
	edgeCount := make(map[int]int)
	for _, edge := range prerequisites {
		from := edge[1]
		to := edge[0]
		nodes[from] = append(nodes[from], to)
		edgeCount[to]++
	}

	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		count := edgeCount[i]
		if count == 0 {
			queue = append(queue, i)
		}
	}

	result := make([]int, 0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)
		for _, adj := range nodes[node] {
			edgeCount[adj]--
			if edgeCount[adj] == 0 {
				queue = append(queue, adj)
				delete(edgeCount, adj)
			}
		}
	}
	if len(edgeCount) > 0 {
		return []int{}
	}
	return result
}

func main() {
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	fmt.Println(findOrder(3, [][]int{{1, 0}, {1, 2}, {0, 1}}))
}
