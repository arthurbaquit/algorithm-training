package main

func containsDuplicate(nums []int) bool {
	contains := make(map[int]bool)
	for _, value := range nums {
		if _, ok := contains[value]; !ok {
			contains[value] = true
			continue
		}
		return true
	}
	return false
}

func main() {
	println(containsDuplicate([]int{1, 2, 3, 1}))
	println(containsDuplicate([]int{1, 2, 3, 4}))
	println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
