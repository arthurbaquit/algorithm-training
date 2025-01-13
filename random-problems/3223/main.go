package main

import "fmt"

func minimumLength(s string) int {
	countL := make([]int, 26)
	for _, b := range s {
		countL[b-'a']++
	}
	ans := 0
	for _, v := range countL {
		if v == 0 {
			continue
		}
		if v%2 == 0 {
			ans += 2
			continue
		}
		ans++
	}
	return ans
}

func main() {
	fmt.Println(minimumLength("abaacbcbb")) // 5
	fmt.Println(minimumLength("aa"))        // 2
}
