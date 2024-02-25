package main

func minimumSemesters(n int, relations [][]int) int {
	coursesRoadmap := make([][]int, n+1)
	needsPre := make([]int, n+1)
	for _, rel := range relations {
		coursesRoadmap[rel[0]] = append(coursesRoadmap[rel[0]], rel[1])
		needsPre[rel[1]]++
	}

	coursesTaken := 0
	sems := 0
	coursesToTake := []int{}
	for i := 1; i < n+1; i++ {
		if needsPre[i] == 0 {
			coursesToTake = append(coursesToTake, i)
		}
	}
	for len(coursesToTake) > 0 {
		sems++
		coursesTaken += len(coursesToTake)
		nextCoursesToTake := []int{}
		for _, each := range coursesToTake {
			for _, to := range coursesRoadmap[each] {
				needsPre[to]--
				if needsPre[to] == 0 {
					nextCoursesToTake = append(nextCoursesToTake, to)
				}
			}
		}
		coursesToTake = nextCoursesToTake

	}

	if coursesTaken < n {
		return -1
	}
	return sems
}

func main() {
	println(minimumSemesters(3, [][]int{{1, 3}, {2, 3}}))
	println(minimumSemesters(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))
	println(minimumSemesters(3, [][]int{{1, 2}, {2, 3}}))
}
