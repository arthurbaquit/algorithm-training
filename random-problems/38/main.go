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

func countAndSayIterative(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}
	ans := ""
	curr := "11"
	for i := 2; i < n; i++ {
		slow, fast, count := 0, 1, 1
		for slow < len(curr) {
			for fast < len(curr) && curr[fast] == curr[slow] {
				fast++
				count++
			}
			countToStr := strconv.Itoa(count)
			ans += countToStr + string(curr[slow])
			slow = fast
			fast = slow + 1
			count = 1
		}
		curr = ans
		ans = ""
	}
	return curr
}

func main() {
	println(countAndSay(5))          //  "111221"
	println(countAndSay(4))          // "1211"
	println(countAndSay(1))          // "1"
	println(countAndSayIterative(5)) // "111221"
	println(countAndSayIterative(4)) // "1211"
	println(countAndSayIterative(1)) // "1"

}
