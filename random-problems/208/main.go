package main

type trieNode struct {
	children map[byte]*trieNode
	isEnd    bool
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: make(map[byte]*trieNode),
		isEnd:    false,
	}
}

type Trie struct {
	root *trieNode
}

func Constructor() Trie {
	return Trie{
		root: newTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	curr := t.root
	for i := range word {
		if _, ok := curr.children[word[i]]; !ok {
			curr.children[word[i]] = newTrieNode()
		}

		curr = curr.children[word[i]]
	}

	curr.isEnd = true
}

func (t *Trie) Search(word string) bool {
	curr := t.root

	for i := range word {
		if _, ok := curr.children[word[i]]; !ok {
			return false
		}
		curr = curr.children[word[i]]
	}
	return curr.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t.root

	for i := range prefix {
		if _, ok := curr.children[prefix[i]]; !ok {
			return false
		}
		curr = curr.children[prefix[i]]
	}

	return true
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	println(trie.Search("apple"))   // returns true
	println(trie.Search("app"))     // returns false
	println(trie.StartsWith("app")) // returns true
	trie.Insert("app")
	println(trie.Search("app")) // returns true
}
