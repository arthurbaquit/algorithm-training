package main

import "fmt"

func moveZeroes(nums []int)  {
    l := len(nums)
    initial, next := 0, 0
    for next < l {
        for next < l && nums[next] == 0 { next++ }
        if next >= l { break }
        nums[initial] = nums[next]
        initial++
        next++
    }
    for i := initial; i<l; i++ {nums[i] = 0}
}

func moveZeroesOpt(nums []int)  {
    l := len(nums)
    initial, next := 0, 0
    for next < l {
        for next < l && nums[next] == 0 { next++ }
        if next >= l { break }
		aux := nums[next]
		nums[next] = nums[initial]
        nums[initial] = aux
        initial++
        next++
    }
}
func main(){
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
	nums = []int{0,1,0,3,12}
	moveZeroesOpt(nums)
	fmt.Println(nums)
	nums = []int{0,0,1}
	moveZeroes(nums)
	fmt.Println(nums)
	nums = []int{0,0,1}
	moveZeroesOpt(nums)
	fmt.Println(nums)
	nums = []int{1,0,0}
	moveZeroes(nums)
	fmt.Println(nums)
	nums = []int{1,0,0}
	moveZeroesOpt(nums)
	fmt.Println(nums)
	nums = []int{0,0,0}
	moveZeroes(nums)
	fmt.Println(nums)
	nums = []int{0,0,0}
	moveZeroesOpt(nums)
	fmt.Println(nums)
}
