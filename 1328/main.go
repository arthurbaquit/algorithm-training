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
	var change byte
	for i := 0; i < 23; i++ {
		change = 'a' + byte(i)
		for j := 0; j < length; j++ {
			if palindrome[j] > change && j != length/2 {
				return strings.Replace(palindrome, string(palindrome[j]), string(change), 1)
			}
		}
	}
	aux := palindrome[length-1] + byte(1)

	return palindrome[:length-1] + string(aux)
}

func main() {
	fmt.Println(breakPalindrome("abccba"))
	fmt.Println(breakPalindrome(("a")))
	fmt.Println(breakPalindrome(("aa")))
	fmt.Println(breakPalindrome(("aba")))
}
