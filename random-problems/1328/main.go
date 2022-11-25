package main

import (
	"fmt"
	"strings"
)

func breakPalindrome(palindrome string) string {
	length := len(palindrome)
	if length == 1 {
		return ""
	}
	for j := 0; j < length; j++ {
		if palindrome[j] != 'a' && j != length/2 {
			return strings.Replace(palindrome, string(palindrome[j]), "a", 1)
		}
	}
	return palindrome[:length-1] + "b"
}

func main() {
	fmt.Println(breakPalindrome("abccba"))
	fmt.Println(breakPalindrome(("a")))
	fmt.Println(breakPalindrome(("aa")))
	fmt.Println(breakPalindrome(("aba")))
}
