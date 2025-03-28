package main

import "math"

func checkPowersOfThree(n int) bool {
	val := n
	checked := make(map[int]bool)
	for val > 0 {
		exp := int(math.Log(float64(val)) / math.Log(3))
		if _, ok := checked[exp]; ok {
			return false
		}
		checked[exp] = true
		expFloat := float64(exp)
		val -= int(math.Pow(3, expFloat))
	}
	return true
}

func checkPowersOfThreeEditorial(n int) bool {
	for n > 0 {
		if n%3 == 2 {
			return false
		}
		n /= 3
	}
	return true
}

func main() {
	println(checkPowersOfThree(12))          // true
	println(checkPowersOfThree(91))          // true
	println(checkPowersOfThree(21))          // false
	println(checkPowersOfThreeEditorial(12)) // true
	println(checkPowersOfThreeEditorial(91)) // true
	println(checkPowersOfThreeEditorial(21)) // false
}
