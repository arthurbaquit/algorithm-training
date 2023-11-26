package main

import (
	"fmt"
	"sort"
)

func maxCoins(piles []int) int {
	sort.Ints(piles)
	ans := 0
	i, f := len(piles)/3, len(piles)-1
	for i > 0 {
		ans += piles[f-1]
		i--
		f -= 2
	}
	return ans
}

func main() {
	piles := []int{2, 4, 1, 2, 7, 8}
	fmt.Println(maxCoins(piles))
}
