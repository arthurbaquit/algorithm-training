package main

func numRabbits(answers []int) int {
	count := 0
	openPairs := make(map[int]int)
	for _, ans := range answers {
		if openPairs[ans] == 0 {
			openPairs[ans] = ans
			count += ans + 1
			continue
		}
		openPairs[ans]--
	}
	return count
}

func main() {
	println(numRabbits([]int{1, 1, 2}))    // 5
	println(numRabbits([]int{10, 10, 10})) // 11
}
