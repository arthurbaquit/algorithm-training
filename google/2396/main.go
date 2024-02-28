package main

import "fmt"

func isStrictlyPalindromic(n int) bool {
	for i := 2; i < n-1; i++ {
		newN := toBaseN(n, i)
		if !checkPal(newN) {
			return false
		}
	}
	return true
}

func toBaseN(n, b int) string {
	res := ""
	for n > 0 {
		res = fmt.Sprintf("%v%v", n%b, res)
		n = (n - n%b) / b
	}
	return res
}
func checkPal(n string) bool {
	for i := 0; i < len(n)/2; i++ {
		if n[i] != n[len(n)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	println(isStrictlyPalindromic(585))
	println(isStrictlyPalindromic(10))
	println(isStrictlyPalindromic(42))
}
