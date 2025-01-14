package main

import "fmt"

func findThePrefixCommonArrayBF(A []int, B []int) []int {
	hasA := make(map[int]bool)
	hasB := make(map[int]bool)
	common := 0
	ans := []int{}
	for i, v := range A {
		hasB[B[i]] = true
		hasA[v] = true
		for key, value := range hasA {
			if value == hasB[key] {
				common++
				delete(hasA, key)
				delete(hasB, key)
			}
		}
		ans = append(ans, common)
	}
	return ans
}

// this only works because A[i] and B[i] <= n, where n <= 50
func findThePrefixCommonArray(A []int, B []int) []int {
	freq := make([]int, len(A)+1)
	common := 0
	ans := []int{}
	for i := range A {
		freq[A[i]]++
		freq[B[i]]++
		if freq[A[i]] == 2 {
			common++
		}
		if A[i] != B[i] && freq[B[i]] == 2 {
			common++
		}
		ans = append(ans, common)
	}
	return ans
}

func main() {
	Ai := []int{1, 3, 2, 4}
	Bi := []int{3, 1, 2, 4}
	fmt.Println(findThePrefixCommonArrayBF(Ai, Bi)) // [0,2,3,4]
	fmt.Println(findThePrefixCommonArray(Ai, Bi))   // [0,2,3,4]
	Ai = []int{2, 3, 1}
	Bi = []int{3, 2, 1}
	fmt.Println(findThePrefixCommonArrayBF(Ai, Bi)) // [0,1,3]
	fmt.Println(findThePrefixCommonArray(Ai, Bi))   // [0,1,3]
}
