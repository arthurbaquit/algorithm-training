package main

import "fmt"

func removeDuplicates(s string) string {
	stack := []rune{}
	for _, r := range s {
		l := len(stack)
		if l > 0 && r == stack[l-1] {
			stack = stack[:l-1]
			continue
		}
		stack = append(stack, r)
	}
	res := ""
	for _, r := range stack {
		res = fmt.Sprintf("%v%v", res, string(r))
	}
	return res
}

func main() {
	fmt.Println(removeDuplicates("abbaca"))
	fmt.Println(removeDuplicates("aaaaaaaa"))
	fmt.Println(removeDuplicates("aaaaaaa"))
}
