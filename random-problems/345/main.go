package main

import "strings"

func reverseVowels(s string) string {
	vowels := "AEIOUaeiou"
	l := len(s)
	first, last := 0, l-1
	res := []byte(s)
	for first < last {
		for first <= last && !strings.Contains(vowels, string(s[first])) {
			first++
		}
		for last > first && !strings.Contains(vowels, string(s[last])) {
			last--
		}
		if first < last {
			res[first], res[last] = res[last], res[first]

		}
		first++
		last--
	}
	return string(res)
}

func main() {
	s := "hello"
	println(reverseVowels(s))
	s = "leetcode"
	println(reverseVowels(s))
}
