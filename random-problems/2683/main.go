package main

// derived[n-1] = original[n-1] ^ original[0]
// derived[n-2] = original[n-2] ^ original[n-1]
// derived[n-3] = original[n-3] ^ original[n-2]
// derived[n-4] = original[n-4] ^ original[n-3]
// ...
// derived[1] = original[1] ^ original[2]
// derived[0] = original[0] ^ original[1]

// (XOR every equation)
// => derived[0] ^ derived[1] ^ ... ^ derived[n-1] = 0

func doesValidArrayExist(derived []int) bool {
	ans := 0
	for _, d := range derived {
		ans ^= d
	}
	return ans == 0
}

func main() {
	println(doesValidArrayExist([]int{1, 1, 0})) // true
	println(doesValidArrayExist([]int{1, 1}))    // true
	println(doesValidArrayExist([]int{1, 0}))    // false
}
