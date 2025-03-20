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

func canPermutePalindromeSinglePass(s string) bool {
	letterCount := make(map[rune]int)
	count := 0
	for _, b := range s {
		letterCount[b]++
		if letterCount[b]%2 == 0 {
			count--
			continue
		}
		count++
	}
	return count <= 1
}

func main() {
	println(canPermutePalindrome("code"))              // false
	println(canPermutePalindrome("aab"))               // true
	println(canPermutePalindrome("carerac"))           // true
	println(canPermutePalindromeSinglePass("code"))    // false
	println(canPermutePalindromeSinglePass("aab"))     // true
	println(canPermutePalindromeSinglePass("carerac")) // true
}
