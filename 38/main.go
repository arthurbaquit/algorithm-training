package main

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	num := countAndSay(n - 1)
	slow, fast := 0, 0
	res := ""
	for fast < len(num)-1 {
		if num[fast] == num[fast+1] {
			fast++
			continue
		}
		res += strconv.Itoa(fast-slow+1) + string(num[slow])
		fast++
		slow = fast
	}
	res += strconv.Itoa(fast-slow+1) + string(num[slow])
	return res
}

func main() {
	println(countAndSay(5))
	println(countAndSay(4))
	println(countAndSay(1))

}
