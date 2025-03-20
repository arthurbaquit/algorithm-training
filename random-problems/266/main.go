package main

func canPermutePalindrome(s string) bool {
	letterCount := make(map[rune]int)
	for _, b := range s {
		letterCount[b]++
	}
	isOdd := false
	for _, val := range letterCount {
		if val%2 == 1 {
			if isOdd {
				return false
			}
			isOdd = true
		}
	}
	return true
}

func main() {
	println(canPermutePalindrome("code"))    // false
	println(canPermutePalindrome("aab"))     // true
	println(canPermutePalindrome("carerac")) // true
}
