package main

import "fmt"

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}
func vowelStrings(words []string, queries [][]int) []int {
	validWordsUntilNow := make([]int, len(words))
	count := 0
	for i, w := range words {
		if isVowel(w[0]) && isVowel(w[len(w)-1]) {
			count++
		}
		validWordsUntilNow[i] = count
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		l, r := query[0], query[1]
		if l == 0 {
			ans[i] = validWordsUntilNow[r]
			continue
		}
		ans[i] = validWordsUntilNow[r] - validWordsUntilNow[l-1]
	}
	return ans
}

func main() {
	fmt.Println(vowelStrings([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}})) // [2,3,0]
	fmt.Println(vowelStrings([]string{"a", "e", "i"}, [][]int{{0, 2}, {0, 1}, {2, 2}}))                  // [3,2,1]
}
