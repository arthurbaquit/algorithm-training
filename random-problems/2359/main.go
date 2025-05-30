package main

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	visitedByNode1 := make(map[int]int)
	visitedByNode2 := make(map[int]int)
	curr := node1
	dist := 0
	for curr != -1 {
		if _, ok := visitedByNode1[curr]; ok {
			// found cycle
			break
		}
		visitedByNode1[curr] = dist
		dist++
		curr = edges[curr]
	}
	curr = node2
	dist = 0
	for curr != -1 {
		if _, ok := visitedByNode2[curr]; ok {
			// found cycle
			break
		}
		visitedByNode2[curr] = dist
		dist++
		curr = edges[curr]
	}
	d, ans := 100001, -1
	for key, value := range visitedByNode1 {
		v, ok := visitedByNode2[key]
		if !ok {
			continue
		}
		maxD := max(value, v)
		if maxD <= d {
			if maxD == d {
				ans = min(ans, key)
				continue
			}
			d = maxD
			ans = key
		}
	}
	return ans
}

func main() {
	println(closestMeetingNode([]int{2, 2, 3, -1}, 0, 1)) // 2
	println(closestMeetingNode([]int{1, 2, -1}, 0, 2))    // 2
}
