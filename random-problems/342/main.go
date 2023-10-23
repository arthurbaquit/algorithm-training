package main

import (
	"fmt"
	"math"
)

func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}
	res := math.Log(float64(n)) / math.Log(4.0)
	return res == math.Trunc(res)
}

func main() {
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(5))
	fmt.Println(isPowerOfFour(1))
}
