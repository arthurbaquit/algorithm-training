package main

func numberOfAlternatingGroups(colors []int, k int) int {
	for i := range k - 1 {
		colors = append(colors, colors[i])
	}
	l, r := 0, 1
	ans := 0
	for r < len(colors) {
		if colors[r] == colors[r-1] {
			l = r
			r++
			continue
		}
		r++
		if r-l < k {
			continue
		}
		ans++
		l++
	}
	return ans
}

func main() {
	println(numberOfAlternatingGroups([]int{0, 1, 0, 1, 0}, 3))       // 3
	println(numberOfAlternatingGroups([]int{0, 1, 0, 0, 1, 0, 1}, 6)) // 2
	println(numberOfAlternatingGroups([]int{1, 1, 0, 1}, 4))          // 0
}
