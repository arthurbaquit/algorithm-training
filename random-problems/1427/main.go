package main

func stringShift(s string, shift [][]int) string {
	dir := 0
	for _, shif := range shift {
		if shif[0] == 0 {
			dir -= shif[1]
			continue
		}
		dir += shif[1]
	}
	dir %= len(s)
	if dir > 0 {
		r := len(s) - dir
		return s[r:] + s[:r]
	}
	l := -dir
	return s[l:] + s[:l]
}

func main() {
	println(stringShift("abc", [][]int{{0, 1}, {1, 2}}))                     // "cab"
	println(stringShift("abcdefg", [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}})) // "efgabcd"
}
