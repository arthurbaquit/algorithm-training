package main

import "fmt"

func findArray(pref []int) []int {

	for i := len(pref) - 1; i >= 1; i-- {
		pref[i] = pref[i] ^ pref[i-1]
	}
	return pref
}

func main() {
	fmt.Println(findArray([]int{1, 2, 3, 4}))
	fmt.Println(([]int{13}))
}
