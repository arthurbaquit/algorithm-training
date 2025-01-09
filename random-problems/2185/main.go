package main

import (
	"fmt"
	"strings"
)

func prefixCount(words []string, pref string) int {
	ans := 0
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			ans++
		}
	}
	return ans
}

/*
Trie version is better if we need to query multiple times for different
prefixes. The time complexity for the trie version is O(n*l) for the
initialization and O(m) for each query, where n is the number of words,
l is the length of the longest word and m is the length of the prefix.
Therefore, once the Trie is initialized, the time complexity for each
individual query is O(m). However yhe time complexity for the brute
force version is O(n*m), where n is the number of words and m is the
length of the prefix.
*/
type Trie struct {
	Frequency int
	Children  map[rune]*Trie
}

func insertWord(root *Trie, word string) {
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

func initializeTrie(words []string) *Trie {
	root := &Trie{
		Frequency: 0,
		Children:  make(map[rune]*Trie),
	}
	for _, word := range words {
		insertWord(root, word)
	}
	return root
}

func prefixCountTrie(root *Trie, pref string) int {
	curr := root
	for _, b := range pref {
		if _, ok := curr.Children[b]; !ok {
			return 0
		}
		curr = curr.Children[b]
	}
	return curr.Frequency
}

func main() {
	words := []string{"apple", "app", "apricot", "april"}
	pref := "ap"
	fmt.Println(prefixCount(words, pref)) // 4
	root := initializeTrie(words)
	fmt.Println(prefixCountTrie(root, pref)) // 4
	words = []string{"pay", "attention", "practice", "attend"}
	pref = "at"
	fmt.Println(prefixCount(words, pref)) // 2
	root = initializeTrie(words)
	fmt.Println(prefixCountTrie(root, pref)) // 2
	words = []string{"leetcode", "win", "loops", "success"}
	pref = "code"
	fmt.Println(prefixCount(words, pref)) // 0
	root = initializeTrie(words)
	fmt.Println(prefixCountTrie(root, pref)) // 0
}
