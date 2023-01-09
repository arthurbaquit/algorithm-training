package main

import "fmt"

func rob(nums []int) int {
    robbedHouses := make(map[int]int)
    var robFrom func(int)int
    robFrom = func(index int) int {
        if index >= len(nums) { return 0 }
        if v, ok := robbedHouses[index]; ok { return v}
        ans := max(robFrom(index+1), robFrom(index+2)+ nums[index]) 
        robbedHouses[index] = ans
        return ans
    }
    return robFrom(0)
}


func max(x,y int) int {
    if x>y {return x}
    return y
}

func main(){
	fmt.Println(rob([]int{1,2,3,1}))
	fmt.Println(rob([]int{2,7,9,3,1}))
	fmt.Println(rob([]int{2,1,1,2}))
}