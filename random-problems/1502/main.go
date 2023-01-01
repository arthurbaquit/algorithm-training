package main

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
    l := len(arr)
    if l <3 {return true}
    sort.Ints(arr)
    for i:=0; i<l-2;i++{
        if arr[i+1] - arr[i] == arr[i+2] - arr[i+1] {continue}
        return false
    }
    return true
}

func main(){
	println(canMakeArithmeticProgression([]int{3,5,1}))
	println(canMakeArithmeticProgression([]int{1,2,4}))
}