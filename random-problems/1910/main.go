package main

func removeOccurrences(s string, part string) string {
	if len(part) > len(s) {
		return s
	}
	l, r := 0, 0
	for l < len(s) && l >= 0 {
		isSub := false
		for r = 0; r < len(part); r++ {
			if l+r >= len(s) || s[l+r] != part[r] {
				isSub = false
				break
			}
			isSub = true
		}
		if !isSub {
			l++
			continue
		}
		s = s[:l] + s[l+r:]
		l -= len(part)
		if l < 0 {
			l = 0
		}
	}
	return s
}

func main() {
	println(removeOccurrences("daabcbaabcbc", "abc")) // dab
	println(removeOccurrences("axxxxyyyyb", "xy"))    // ab
}
