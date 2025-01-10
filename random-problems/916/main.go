package main

import "fmt"

func wordSubsets(words1 []string, words2 []string) []string {
	ans := []string{}
	allW2Letters := make(map[rune]int)
	for _, w2 := range words2 {
		w2L := make(map[rune]int)
		for _, w := range w2 {
			w2L[w]++
		}
		for key, value := range w2L {
			if allW2Letters[key] >= value {
				continue
			}
			allW2Letters[key] = value
		}
	}
	for _, w1 := range words1 {
		w1L := make(map[rune]int)
		isUniversal := true
		for _, w := range w1 {
			w1L[w]++
		}
		for key, value := range allW2Letters {
			if w1L[key] >= value {
				continue
			}
			isUniversal = false
			break
		}
		if isUniversal {
			ans = append(ans, w1)
		}
	}
	return ans
}

func main() {
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"})) // ["facebook","google","leetcode"]
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"l", "e"})) // ["apple","google","leetcode"]
}
