package main

import "fmt"

func combine(n int, k int) [][]int {
    var res [][]int

    combination := make([]int, k)
    
    var backtrack func(int)
    backtrack = func(index int) {
        if index == k {
            copiedCombination := make([]int, k)
            copy(copiedCombination, combination)
            res = append(res, copiedCombination)
            return 
        }
        
        begin := 1
        
        if index > 0 {
            begin = combination[index-1] + 1
        }
        
        for i := begin; i <= n; i++ {
            combination[index] = i
            backtrack(index + 1)
        } 
    }
    
    backtrack(0)
    
    return res
}

func main(){
	fmt.Println(combine(4, 2))
}