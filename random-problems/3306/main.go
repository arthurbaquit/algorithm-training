package main

func countOfSubstrings(word string, k int) int64 {
	return countOfSubstringsAtLeastK(word, k) - countOfSubstringsAtLeastK(word, k+1)
}
func countOfSubstringsAtLeastK(word string, k int) int64 {
	countLetter := []int{0, 0, 0, 0, 0, 0}
	l, r := 0, 0
	var ans int64
	for r < len(word) {
		countLetter[getIdx(word[r])]++
		for countLetter[0] > 0 && countLetter[1] > 0 && countLetter[2] > 0 && countLetter[3] > 0 && countLetter[4] > 0 && countLetter[5] >= k {
			ans += int64(len(word) - r)
			countLetter[getIdx(word[l])]--
			l++
		}
		r++
	}
	return ans
}
func getIdx(b byte) int {
	switch b {
	case 'a':
		return 0
	case 'e':
		return 1
	case 'i':
		return 2
	case 'o':
		return 3
	case 'u':
		return 4
	default:
		return 5
	}
}

func main() {
	println(countOfSubstrings("aeioqq", 1))         // 0
	println(countOfSubstrings("aeiou", 0))          // 1
	println(countOfSubstrings("ieaouqqieaouqq", 1)) // 3
}
