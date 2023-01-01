package main

func arraySign(nums []int) int {
    prod := 1
    for _, num := range(nums) {
        prod *= signFunc(num)
        if prod == 0 { return 0}
    }
    return prod
}

func signFunc(x int) int {
    if x == 0 {return 0}
    if x>0 {return 1}
    return -1
}

func main(){
	println(arraySign([]int{-1,-2,-3,-4,3,2,1}))
	println(arraySign([]int{1,5,0,2,-3}))
	println(arraySign([]int{-1,1,-1,1,-1}))
}