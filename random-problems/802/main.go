package main

import "fmt"

func eventualSafeNodes(graph [][]int) []int {
	ans := make([]int, 0, len(graph))
	var safeNodes [10000]bool
	var visited [10000]bool
	for i := range graph {
		if safeNodes[i] || isSafe(graph, i, &visited, &safeNodes) {
			ans = append(ans, i)
		}
	}

	return ans
}

func isSafe(graph [][]int, i int, visited, safeNodes *[10000]bool) bool {
	if safeNodes[i] {
		return true
	}
	// if I already visited and it's not a safeNode, then it has a cycle
	// then I can return false
	if visited[i] {
		return false
	}
	visited[i] = true

	// check if any of its child has cycle
	for _, child := range graph[i] {
		if !isSafe(graph, child, visited, safeNodes) {
			return false
		}
	}

	// if it doesn't, return true and mark it as safeNode for
	// avoiding checking it again
	safeNodes[i] = true
	return safeNodes[i]
}

func main() {
	fmt.Println(eventualSafeNodes([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}))    // [2,4,5,6]
	fmt.Println(eventualSafeNodes([][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}})) // [4]
}
