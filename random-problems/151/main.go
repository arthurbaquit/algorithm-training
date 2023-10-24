package main

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	first, last := 0, len(words)-1
	for first < last {
		for first < last && words[first] == "" {
			first++
		}
		for last > first && words[last] == "" {
			last--
		}
		if first < last {
			words[first], words[last] = words[last], words[first]
			first++
			last--
		}
	}
	res := []string{}
	for _, v := range words {
		if v != "" {
			res = append(res, v)
		}
	}
	return strings.Join(res, " ")
}

func main() {
	s := "the sky is blue"
	println(reverseWords(s))
	s = "  hello world!  "
	println(reverseWords(s))
	s = "a good   example"
	println(reverseWords(s))
	s = "  Bob    Loves  Alice   "
	println(reverseWords(s))
	s = "Alice does not even like bob"
	println(reverseWords(s))
}
