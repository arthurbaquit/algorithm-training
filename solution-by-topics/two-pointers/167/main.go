package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	l,r := 0, len(numbers)-1
	 for l<r {
		 aux := numbers[l] + numbers[r]
		 if aux == target { return []int{l+1,r+1}}
		 if aux > target { r--; continue}
		 l++
	 }
	 return []int{}
 }

func main(){
	fmt.Println(twoSum([]int{2,7,11,15},9))
	fmt.Println(twoSum([]int{2,3,4},6))
	fmt.Println(twoSum([]int{-1,0},-1))
}