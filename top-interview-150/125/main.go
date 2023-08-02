package main

import "fmt"

func isAlphaNumeric(b byte) bool {
	return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9')
}

func isPalindrome(s string) bool {
	begin, end := 0, len(s)-1

	for begin < end {
		for begin < end && !isAlphaNumeric(s[begin]) {
			begin++
		}
		for end > begin && !isAlphaNumeric(s[end]) {
			end--
		}
		if begin > end {
			return false
		}
		beginL, endL := s[begin], s[end]
		if 'A' <= beginL && beginL <= 'Z' {
			beginL = beginL + 32
		}
		if 'A' <= endL && endL <= 'Z' {
			endL = endL + 32
		}
		if beginL != endL {
			return false
		}
		begin++
		end--
	}
	return true
}

func main() {

	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome(".,"))
	fmt.Println(isPalindrome("0P"))
}
