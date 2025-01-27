package main

import "fmt"

func isPrereq(preReqs map[int][]int, visited []bool, src, target int) bool {
	visited[src] = true
	if src == target {
		return true
	}
	for _, pr := range preReqs[src] {
		if !visited[pr] {
			if isPrereq(preReqs, visited, pr, target) {
				return true
			}
		}
	}
	return false
}
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	preReqs := make(map[int][]int)
	for _, preReq := range prerequisites {
		preReqs[preReq[0]] = append(preReqs[preReq[0]], preReq[1])
	}
	ans := []bool{}
	for _, querie := range queries {
		visited := make([]bool, numCourses)
		ans = append(ans, isPrereq(preReqs, visited, querie[0], querie[1]))
	}
	return ans
}

func main() {
	fmt.Println(checkIfPrerequisite(2, [][]int{{1, 0}}, [][]int{{0, 1}, {1, 0}}))                 // [false, true]
	fmt.Println(checkIfPrerequisite(2, [][]int{}, [][]int{{0, 1}, {1, 0}}))                       // [false, false]
	fmt.Println(checkIfPrerequisite(3, [][]int{{1, 2}, {1, 0}, {2, 0}}, [][]int{{1, 0}, {1, 2}})) // [true, true]
}
