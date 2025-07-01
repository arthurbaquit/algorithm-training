package main

func possibleStringCount(word string) int {
	ans := 1
	for i := 1; i < len(word); i++ {
		lastByte := word[i-1] - 'a'
		currByte := word[i] - 'a'
		if currByte == lastByte {
			ans++
		}
	}
	return ans
}

func main() {
	println(possibleStringCount("abbcccc")) // 5
	println(possibleStringCount("abcd"))    // 1
	println(possibleStringCount("aaaa"))    // 4
}
