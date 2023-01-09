package main

import "fmt"

func permute(nums []int) [][]int {
    var res [][]int
    n := len(nums)
    permutations := make([]int, n)
    var backtrack func(int)
    backtrack = func(index int){
        if index == n{
            copiedPermutations := make([]int, n)
            copy(copiedPermutations, permutations)
            res = append(res, copiedPermutations)
            return
        }
        for i:=index; i<n; i++{
            permutations[index] = nums[i]
            nums[index], nums[i] = nums[i], nums[index]
            backtrack(index + 1)
            nums[index], nums[i] = nums[i], nums[index]
        }
    }
    backtrack(0)
    return res
}

func main(){
	fmt.Println(permute([]int{1,2,3}))
}