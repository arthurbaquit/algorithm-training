package main

import "fmt"

func garbageCollection(garbage []string, travel []int) int {
	paperMin := 0
	metalMin := 0
	glassMin := 0
	for i := len(garbage) - 1; i > -1; i-- {
		// glassMin+=strings.Count(garbage[i], "G")
		// metalMin+=strings.Count(garbage[i], "M")
		// paperMin+=strings.Count(garbage[i], "P")
		for _, g := range garbage[i] {
			if g == 'G' {
				glassMin++
				continue
			}
			if g == 'M' {
				metalMin++
				continue
			}
			paperMin++
		}
		if i == 0 {
			break
		}
		if glassMin > 0 {
			glassMin += travel[i-1]
		}
		if metalMin > 0 {
			metalMin += travel[i-1]
		}
		if paperMin > 0 {
			paperMin += travel[i-1]
		}
	}
	return glassMin + metalMin + paperMin
}

func main() {
	garbage := []string{"MPPMMPM", "PMMGGP", "MPMCG"}
	travel := []int{3, 6, 9}
	fmt.Println(garbageCollection(garbage, travel))
	garbage = []string{"G", "P", "GP", "GG"}
	travel = []int{2, 4, 3}
	fmt.Println(garbageCollection(garbage, travel))
}
