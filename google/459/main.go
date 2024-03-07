package main

import "strings"

func repeatedSubstringPattern(s string) bool {
	return strings.Contains(s[1:]+s[0:len(s)-1], s)
}

func main() {
	println(repeatedSubstringPattern("abab"))
	println(repeatedSubstringPattern("aba"))
	println(repeatedSubstringPattern("abcabcabcabc"))
}
