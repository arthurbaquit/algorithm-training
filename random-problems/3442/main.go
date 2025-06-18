package main

func maxDifference(s string) int {
	charFreq := make([]int, 26)
	for i := range s {
		charFreq[s[i]-'a']++
	}
	maxOdd, minEven := 0, 1000
	for _, freq := range charFreq {
		if freq == 0 {
			continue
		}
		if freq%2 == 0 {
			minEven = min(minEven, freq)
			continue
		}
		maxOdd = max(maxOdd, freq)
	}
	return maxOdd - minEven
}

func main() {
	println(maxDifference("aaaaabbc")) // 3
	println(maxDifference("abcabcab")) // 1
}
