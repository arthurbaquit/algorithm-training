package main

func numberOfSubstrings(s string) int {
	countLetter := []int{0, 0, 0}
	l, r, count := 0, 0, 0
	for r < len(s) {
		countLetter[s[r]-'a'] += 1
		for countLetter[0] > 0 && countLetter[1] > 0 && countLetter[2] > 0 {
			count += len(s) - r
			countLetter[s[l]-'a']--
			l++
		}
		r++
	}
	return count
}

func main() {
	println(numberOfSubstrings("abcabc")) // 10
	println(numberOfSubstrings("aaacb"))  // 3
	println(numberOfSubstrings("abc"))    // 1
}
