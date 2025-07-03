package main

var word []byte

func kthCharacter(k int) byte {
	if len(word) == 0 {
		word = append(word, 'a')
	}
	for len(word) < 500 {
		buf := make([]byte, len(word))
		for i := 0; i < len(word); i++ {
			if word[i] == 'z' {
				buf[i] = 'a'
				continue
			}
			buf[i] = word[i] + 1
		}
		word = append(word, buf...)

	}
	return word[k-1]
}

func main() {
	println(string(kthCharacter(5)))  // b
	println(string(kthCharacter(10))) // c
}
