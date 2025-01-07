package main

import (
	"fmt"
	"strings"
)

func stringMatchingBF(words []string) []string {
	ans := []string{}
	checked := make(map[string]bool)
	for i, word := range words {
		for j, w := range words {
			if i == j {
				continue
			}
			if strings.Contains(word, w) && !checked[w] {
				ans = append(ans, w)
				checked[w] = true
			}
		}
	}
	return ans
}

type Trie struct {
	Frequency int
	Children  map[rune]*Trie
}

func insertSuffix(root *Trie, word string) {
	curr := root
	for _, b := range word {
		if _, ok := curr.Children[b]; !ok {
			curr.Children[b] = &Trie{
				Frequency: 0,
				Children:  make(map[rune]*Trie),
			}
		}
		curr = curr.Children[b]
		curr.Frequency++
	}
}

func isSubstring(root *Trie, word string) bool {
	curr := root
	for _, b := range word {
		// as the word was inserted previously, we have the garantee
		// that we can transverse it until the end
		curr = curr.Children[b]
	}
	return curr.Frequency > 1
}
func stringMatching(words []string) []string {
	ans := []string{}
	root := Trie{
		Frequency: 0,
		Children:  make(map[rune]*Trie),
	}
	for _, word := range words {
		for i := range word {
			insertSuffix(&root, word[i:])
		}
	}
	for _, word := range words {
		if isSubstring(&root, word) {
			ans = append(ans, word)
		}
	}
	return ans
}

func main() {
	words := []string{"mass", "as", "hero", "superhero"}
	fmt.Println(stringMatching(words))   // ["as","hero"]
	fmt.Println(stringMatchingBF(words)) // ["as","hero"]
	words = []string{"leetcoder", "leetcode", "od", "hamlet", "am"}
	fmt.Println(stringMatching(words))   // ["leetcode","od","am"]
	fmt.Println(stringMatchingBF(words)) // ["leetcode","od","am"]
}
