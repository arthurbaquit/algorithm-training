package main

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	prefix := make([]int, n)
	for _, shift := range shifts {
		dir := -1
		if shift[2] == 1 {
			dir = 1
		}
		prefix[shift[0]] += dir
		if shift[1]+1 < n {
			prefix[shift[1]+1] -= dir
		}
	}
	prefix[0] %= 26
	byteS := make([]byte, n)
	for i := 0; i < n; i++ {
		if i > 0 {
			prefix[i] = (prefix[i] + prefix[i-1]) % 26
		}
		val := int(s[i]) + prefix[i]
		if val > 'z' {
			val -= 26
		}
		if val < 'a' {
			val += 26
		}
		byteS[i] = byte(val)
	}
	return string(byteS)
}

func main() {
	println(shiftingLetters("abc", [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}})) // "ace"
	println(shiftingLetters("dztz", [][]int{{0, 0, 0}, {1, 1, 1}}))           // "catz"
}
