package main

import "fmt"

func addBinary(a string, b string) string {
	la, lb := len(a)-1, len(b)-1
	res := ""
	flag := 0
	for la >= 0 || lb >= 0 {
		sum := flag
		if la >= 0 {
			sum += int(a[la] - '0')
			la--
		}
		if lb >= 0 {
			sum += int(b[lb] - '0')
			lb--
		}
		flag = 0
		if sum > 1 {
			flag = 1
		}
		res = fmt.Sprintf("%v%v", sum%2, res)
	}
	if flag > 0 {
		res = "1" + res
	}
	return res
}

func main() {
	a := "11"
	b := "1"
	println(addBinary(a, b))
	a = "1010"
	b = "1011"
	println(addBinary(a, b))
}
