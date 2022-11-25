package main

import "fmt"

func findErrorNums(nums []int) []int {
	l := len(nums)
	res := []int{}
	count := 0
	visited := make([]int, l+1)
	for _, v := range nums {
		count += v
		if visited[v] == 0 {
			visited[v] = v
			continue
		}
		res = append(res, v)
	}
	res = append(res, l*(l+1)/2-count+res[0])
	return res
}

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
	fmt.Println(findErrorNums([]int{1, 1}))
	fmt.Println(findErrorNums([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 20}))
}
