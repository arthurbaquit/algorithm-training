package main

func maximumSum(nums []int) int {
	biggestVal := make([]int, 82)
	ans := -1
	for _, num := range nums {
		sdig := sumDigits(num)
		v := biggestVal[sdig]
		if v == 0 {
			biggestVal[sdig] = num
			continue
		}
		ans = max(ans, v+num)
		biggestVal[sdig] = max(v, num)
	}
	return ans
}

func sumDigits(x int) int {
	sum := 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	return sum
}

func main() {
	println(maximumSum([]int{18, 43, 36, 13, 7})) // 54
	println(maximumSum([]int{10, 12, 19, 14}))    // -1
}
