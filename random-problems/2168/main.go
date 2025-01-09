package main

// The BruteForce takes longer than O(n^2) time complexity.
// because the hashing a string of length n takes O(n) time.
// To fix this, we can try rolling hash. However, to approach
// differently, we can use a Trie data structure.

func equalDigitFrequencyBF(s string) int {
	ans := make(map[string]int)
	for i := 0; i < len(s); i++ {
		check := make(map[byte]int)
		for j := i; j < len(s); j++ {
			curr := s[i : j+1]
			check[s[j]]++
			aux := 0
			for _, v := range check {
				if aux == 0 {
					aux = v
					continue
				}
				if v != aux {
					aux = -1
					break
				}
			}
			if aux != -1 {
				ans[curr]++
			}
		}
	}
	return len(ans)
}

type TrieNode struct {
	IsVisited bool
	Children  map[uint8]*TrieNode
}

func equalDigitFrequency(s string) int {
	ans := 0
	root := &TrieNode{
		IsVisited: false,
		Children:  make(map[uint8]*TrieNode),
	}
	for i := 0; i < len(s); i++ {
		curr := root
		digitFrequency := make([]int, 10)
		uniqueCount := 0
		maxFreq := 0
		for j := i; j < len(s); j++ {
			currDig := s[j] - '0'
			if digitFrequency[currDig] == 0 {
				uniqueCount++
			}
			digitFrequency[currDig]++
			if digitFrequency[currDig] > maxFreq {
				maxFreq = digitFrequency[currDig]
			}
			if _, ok := curr.Children[currDig]; !ok {
				curr.Children[currDig] = &TrieNode{
					IsVisited: false,
					Children:  make(map[uint8]*TrieNode),
				}
			}
			curr = curr.Children[currDig]
			if uniqueCount*maxFreq == j-i+1 && !curr.IsVisited {
				ans++
				curr.IsVisited = true
			}
		}
	}
	return ans
}

func main() {
	println(equalDigitFrequencyBF("1212"))  // 5
	println(equalDigitFrequency("1212"))    // 5
	println(equalDigitFrequencyBF("12321")) // 9
	println(equalDigitFrequency("12321"))   // 9
}
