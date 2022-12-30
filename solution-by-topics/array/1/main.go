package main

import "fmt"

func twoSum(nums []int, target int) []int {
    sol := make(map[int]int)
    for idx, val := range(nums) {
        aux := target-val
        if _, ok := sol[aux]; ok {
            return []int{idx, sol[aux]}
        }
        sol[val] = idx
    }
    return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{3,2,4}, 6))
	fmt.Println(twoSum([]int{3,3}, 6))
}