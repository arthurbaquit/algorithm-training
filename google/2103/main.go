package main

import "strings"

func countPoints(rings string) int {
	res := make(map[byte]string)
	for i := 0; i < len(rings)-1; i = i + 2 {
		if strings.Contains(res[rings[i+1]], string(rings[i])) {
			continue
		}
		res[rings[i+1]] += string(rings[i])
	}
	count := 0
	for _, v := range res {
		if strings.Contains(v, "R") && strings.Contains(v, "G") && strings.Contains(v, "B") {
			count++
		}
	}
	return count
}

func main() {
	println(countPoints("B0B6G0R6R0R6G9"))
	println(countPoints("B0R0G0R9R0B0G0"))
	println(countPoints("G4"))
}
