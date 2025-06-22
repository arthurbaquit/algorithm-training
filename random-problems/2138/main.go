package main

import (
	"fmt"
	"strings"
)

func divideString(s string, k int, fill byte) []string {
	lastChunkSize := len(s) % k
	if lastChunkSize != 0 {
		s += strings.Repeat(string(fill), k-lastChunkSize)
	}
	ans := make([]string, len(s)/k)
	for i := 0; i <= len(s)-k; i += k {
		ans[i/k] = s[i : i+k]
	}
	return ans
}

func main() {
	fmt.Println(divideString("abcdefghi", 3, 'x'))  // ["abc", "def", "ghi"]
	fmt.Println(divideString("abcdefghij", 3, 'x')) // ["abc", "def", "ghi", "jxx"]
}
