package main

func makeFancyString(s string) string {
	if len(s) < 3 {
		return s
	}
	ans := make([]byte, 0, len(s))
	for l, m, r := 0, 1, 2; r < len(s); l, m, r = l+1, m+1, r+1 {
		if s[l] == s[m] && s[m] == s[r] {
			continue
		}
		ans = append(ans, s[l])
	}
	ans = append(ans, s[len(s)-2:]...)
	return string(ans)
}

func main() {
	println(makeFancyString("leeetcode")) // "leetcode"
	println(makeFancyString("aaabaaaa"))  // "aabaa"
	println(makeFancyString("aab"))       // "aab"
}
