package main

func numberOfMatches(n int) int {
	ans := 0
	for n > 1 {
		if n%2 == 0 {
			n /= 2
			ans += n
			continue
		}
		n = (n-1)/2 + 1
		ans += n - 1
	}
	return ans
}

func main() {
	println(numberOfMatches(7))
}
