package main

import (
	"math"
	"strconv"
)

// that's the right approach, but we can minimize the number of fors by working
// with whiles for the 31 bits for a int. However I found it easier to just
// convert the numbers to binary strings and work with them. The time complexity
// is O(log(n)) where n is the maximum number between num1 and num2 regardless
// of the approach.

func minimizeXor(num1 int, num2 int) int {
	num1B := strconv.FormatInt(int64(num1), 2)
	num2B := strconv.FormatInt(int64(num2), 2)
	oneCount := 0
	for _, b := range num2B {
		if b == '1' {
			oneCount++
		}
	}
	ans := 0
	for i, b := range num1B {
		if b == '1' {
			if oneCount > 0 {
				oneCount--
				// We need to take care because the num1B is inverted.
				// Therefore, the index we are checking it's not the
				// same as the index in the binary representation.
				ans += powInt(2, len(num1B)-i-1)
				continue
			}
			break
		}
	}
	idx := len(num1B) - 1
	for oneCount > 0 && idx >= 0 {
		if num1B[idx] == '1' {
			idx--
			continue
		}
		oneCount--
		ans += powInt(2, len(num1B)-idx-1)
		idx--
	}
	for oneCount > 0 {
		i := len(num1B) + oneCount - 1
		ans += powInt(2, i)
		oneCount--
	}
	return ans
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func main() {
	println(minimizeXor(3, 5))   // 3
	println(minimizeXor(1, 12))  // 3
	println(minimizeXor(65, 84)) // 67
}
