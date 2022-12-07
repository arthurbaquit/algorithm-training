package main

import "strings"

func halvesAreAlike(s string) bool {
	vowel := "AEIOUaeiou"
	l := len(s)
	cf, cl := 0, 0
	for i := 0; i < l/2; i++ {
		if strings.Contains(vowel, string(s[i])) {
			cf++
		}
		if strings.Contains(vowel, string(s[l-i-1])) {
			cl++
		}
	}
	return cf == cl
}

func main() {
	println(halvesAreAlike("book"))
	println(halvesAreAlike("textbook"))
	println(halvesAreAlike("MerryChristmas"))
	println(halvesAreAlike("AbCdEfGh"))
}
