package main

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	nums := map[rune]bool{
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
		'0': true,
	}
	hasVowel, hasConsonant := false, false
	for _, r := range word {
		if _, ok := vowels[r]; ok {
			hasVowel = true
			continue
		}
		if _, ok := vowels[r]; !ok {
			if _, ok := nums[r]; !ok {
				if r == '@' || r == '#' || r == '$' {
					return false
				}
				hasConsonant = true
			}
		}
	}
	return hasVowel && hasConsonant
}

func main() {
	println(isValid("234Adas")) // true
	println(isValid("b3"))      // false
	println(isValid("a3$e"))    // false

}
