package main

import "fmt"

func lengthOfLastWord(s string) int {
	count, aux := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			count = 0
			continue
		}
		count++
		aux = count
	}
	return aux
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("Test with spaces in the end   "))
	fmt.Println(lengthOfLastWord("   Test with spaces in the beginning"))
}
