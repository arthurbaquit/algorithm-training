package main

func triangleType(nums []int) string {
	l1, l2, l3 := nums[0], nums[1], nums[2]
	if l1+l2 <= l3 || l1+l3 <= l2 || l3+l2 <= l1 {
		return "none"
	}
	if l1 == l2 && l2 == l3 {
		return "equilateral"
	}
	if l1 != l2 && l1 != l3 && l3 != l2 {
		return "scalene"
	}
	return "isosceles"
}

func main() {
	println(triangleType([]int{3, 3, 3})) // equilateral
	println(triangleType([]int{3, 4, 5})) // scalene
}
