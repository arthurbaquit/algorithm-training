package main

func canConstruct(s string, k int) bool {
	if k > len(s) {
		return false
	}
	if k == len(s) {
		return true
	}
	lettersCount := make([]int, 26)
	for _, l := range s {
		lettersCount[l-'a']++
	}
	odd := 0
	for _, v := range lettersCount {
		if v%2 == 1 {
			odd++
		}
	}
	return odd <= k
}

func main() {
	println(canConstruct("annabelle", 2)) // true
	println(canConstruct("leetcode", 3))  // false
	println(canConstruct("true", 4))      // true
}
