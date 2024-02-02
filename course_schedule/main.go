package main

import "fmt"

/*
There are a total of numCourses courses you have to take, labeled from 0 to
numCourses - 1. You are given an array prerequisites where prerequisites[i] =
[ai, bi] indicates that you must take course bi first if you want to take course
ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first
take course 1.

Return true if you can finish all courses. Otherwise, return false.



Example 1:

Input: numCourses = 2, prerequisites = [[1,0]] Output: true Explanation: There
are a total of 2 courses to take. To take course 1 you should have finished
course 0. So it is possible.

Example 2:

Input: numCourses = 2, prerequisites = [[1,0],[0,1]] Output: false Explanation:
There are a total of 2 courses to take. To take course 1 you should have
finished course 0, and to take course 0 you should also have finished course 1.
So it is impossible.



Constraints:

1 <= numCourses <= 2000 0 <= prerequisites.length <= 5000
prerequisites[i].length == 2 0 <= ai, bi < numCourses All the pairs
prerequisites[i] are unique.

*/

/*
1. check what nodes don't have any incoming connections
2. remove all outgoing connections from the node
3. step 1
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}
	nodes := make(map[int][]int)
	edgeCount := make(map[int]int)

	for _, edges := range prerequisites {
		from := edges[1]
		to := edges[0]
		nodes[from] = append(nodes[from], to)
		edgeCount[to]++
	}

	return re(nodes, edgeCount)
}

func re(nodes map[int][]int, edgeCount map[int]int) bool {
	// fmt.Print("from ", nodes, " ")
	if len(nodes) == 0 {
		return true
	}
	n := -1
	for node, adjs := range nodes {
		val := edgeCount[node]
		if val == 0 && len(adjs) > 0 {
			n = node
			break
		}
	}
	// fmt.Println("selected", n)
	if n == -1 {
		return false
	}
	for _, node := range nodes[n] {
		edgeCount[node]--
	}
	delete(nodes, n)
	return re(nodes, edgeCount)
}

func main() {
	fmt.Println(canFinish(20, [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}})) // false
	fmt.Println(canFinish(2, [][]int{{1, 0}}))                                                                  // true
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))                                                          // false
	fmt.Println(canFinish(2, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}))                                          // true
	fmt.Println(canFinish(2, [][]int{{1, 0}, {1, 2}, {0, 1}}))                                                  // false
}
