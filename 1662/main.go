package main

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	w1, w2 := "", ""
	for i := 0; i < len(word1); i++ {
		w1 += word1[i]
	}
	for i := 0; i < len(word2); i++ {
		w2 += word2[i]
	}
	if len(w1) != len(w2) {
		return false
	}
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			return false
		}
	}
	return true
}

func main() {
	word1 := []string{"ab", "c"}
	word2 := []string{"a", "bc"}
	println(arrayStringsAreEqual(word1, word2))
	word1 = []string{"a", "cb"}
	word2 = []string{"ab", "c"}
	println(arrayStringsAreEqual(word1, word2))
}
