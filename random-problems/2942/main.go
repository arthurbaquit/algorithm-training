package main

import (
	"fmt"
	"strings"
)

func findWordsContaining(words []string, x byte) []int {
	ans := []int{}
	for i, word := range words {
		if strings.IndexByte(word, x) != -1 {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	fmt.Println(findWordsContaining([]string{"leet", "code"}, byte('e')))              // [0, 1]
	fmt.Println(findWordsContaining([]string{"abc", "bcd", "aaaa", "cbc"}, byte('a'))) // [0, 2]
	fmt.Println(findWordsContaining([]string{"abc", "bcd", "aaaa", "cbc"}, byte('z'))) // []
}
