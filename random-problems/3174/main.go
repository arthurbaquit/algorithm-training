package main

func clearDigits(s string) string {
	l := 0
	for l < len(s) && l >= 0 {
		if s[l] >= 'a' && s[l] <= 'z' {
			l++
			continue
		}
		s = s[:l-1] + s[l+1:]
		l = l - 1
	}
	return s
}

func main() {
	println(clearDigits("abc"))  // abc
	println(clearDigits("cb34")) // ""
}
