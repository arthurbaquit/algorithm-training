package main

import "strings"

func removeVowelsByHand(S string) string {
	var ans string
	for _, c := range S {
		if c != 'a' && c != 'e' && c != 'i' && c != 'o' && c != 'u' {
			ans += string(c)
		}
	}
	return ans
}

func removeVowels(s string) string {

	s = strings.ReplaceAll(s, "a", "")
	s = strings.ReplaceAll(s, "e", "")
	s = strings.ReplaceAll(s, "i", "")
	s = strings.ReplaceAll(s, "o", "")
	s = strings.ReplaceAll(s, "u", "")
	return s
}

func main() {
	println(removeVowels("leetcodeisacommunityforcoders"))
	println(removeVowels("aeiou"))
	println(removeVowelsByHand("leetcodeisacommunityforcoders"))
	println(removeVowelsByHand("aeiou"))
}
