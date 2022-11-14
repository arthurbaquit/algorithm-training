package main

func mySqrt(x int) int {
	res := 0
	for (res+1)*(res+1) <= x {
		res++
	}
	return res
}

func mySqrtBinaryTree(x int) int {
	l, m, r := 0, x/2, x
	if x == 0 {
		return 0
	}
	for l <= r {
		m = (r + l) / 2
		if m == 0 {
			return 1
		}
		if x == m*m {
			return m
		}
		if x > m*m {
			if x < (m+1)*(m+1) {
				return m
			}
			l = m + 1
			continue
		}
		r = m - 1
	}
	return m
}

func main() {
	println(mySqrt(8))
	println(mySqrtBinaryTree(8))
	println(mySqrt(0))
	println(mySqrtBinaryTree(0))
	println(mySqrt(1))
	println(mySqrtBinaryTree(1))
	println(mySqrt(80))
	println(mySqrtBinaryTree(80))
}
