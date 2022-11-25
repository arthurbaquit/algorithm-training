package main

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	evenSum := 0
	result := make([]int, len(queries))
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			evenSum += nums[i]
		}
	}
	for i := 0; i < len(queries); i++ {
		if nums[queries[i][1]]%2 == 0 {
			evenSum -= nums[queries[i][1]]
		}
		nums[queries[i][1]] += queries[i][0]
		if nums[queries[i][1]]%2 == 0 {
			evenSum += nums[queries[i][1]]
		}
		result[i] = evenSum
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	queries := [][]int{
		{1, 0}, {-3, 1}, {-4, 0}, {2, 3},
	}
	result := sumEvenAfterQueries(nums, queries)
	for i := 0; i < len(result); i++ {
		print(result[i], " ")
	}
}
