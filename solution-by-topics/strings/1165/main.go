package main

func calculateTime(keyboard string, word string) int {
	count := 0
	previous := 0
	kb := make(map[rune]int)
	for i, v := range keyboard {
		kb[v] = i
	}
	for _, v := range word {
		count += Abs(previous, kb[v])
		previous = kb[v]
	}
	return count
}

func Abs(x, y int) int {
	if x-y < 0 {
		return y - x
	}
	return x - y
}

func main() {
	println(calculateTime("pqrstuvwxyzabcdefghijklmno", "leetcode"))
	println(calculateTime("abcdefghijklmnopqrstuvwxyz", "cba"))
}
