package main

func guess(num int) int {
	if num == 6 {
		return 0
	}
	if num > 6 {
		return -1
	}
	return 1
}

func guessNumber(n int) int {
	l, m, r := 0, n/2, n
	for l <= r {
		m = (l + r) / 2
		g := guess(m)
		if g == -1 {
			r = m - 1
			continue
		}
		if g == 1 {
			l = m + 1
			continue
		}
		break
	}
	return m

}

func main() {
	println(guessNumber(10))
	println(guessNumber(20))
	println(guessNumber(30))

}
