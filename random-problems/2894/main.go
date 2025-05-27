package main

func differenceOfSums(n int, m int) int {
	sumFrom1ToN := n * (n + 1) / 2
	a1m := m
	anm := (n - n%m)
	smToN := (a1m + anm) * (anm) / (2 * m)
	return sumFrom1ToN - 2*smToN
}

func main() {
	println(differenceOfSums(10, 3)) // 19
	println(differenceOfSums(5, 6))  // 15
	println(differenceOfSums(5, 1))  // -15
}
