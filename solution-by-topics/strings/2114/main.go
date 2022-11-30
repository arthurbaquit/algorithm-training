package main

func mostWordsFound(sentences []string) int {
	max := 0
	for _, sentence := range sentences {
		count := 1
		for i := 0; i < len(sentence); i++ {
			if sentence[i] == ' ' {
				count++
			}
		}
		max = Max(max, count)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	println(mostWordsFound([]string{"hello world", "this is a sentence", "hello hello hello hello"}))
	println(mostWordsFound([]string{"one", "two", "three"}))
	println(mostWordsFound([]string{"a", "b", "c", "d", "e"}))
}
