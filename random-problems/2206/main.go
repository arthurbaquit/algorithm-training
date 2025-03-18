package main

func divideArray(nums []int) bool {
	numCount := make(map[int]int)
	for _, num := range nums {
		numCount[num]++
	}
	for _, value := range numCount {
		if value%2 != 0 {
			return false
		}
	}
	return true
}

func main() {
	println(divideArray([]int{1, 2, 3, 3, 2, 1})) // true
	println(divideArray([]int{1, 2, 3, 3, 2, 2})) // false
	println(divideArray([]int{1, 1, 1, 1, 1, 1})) // true
}
