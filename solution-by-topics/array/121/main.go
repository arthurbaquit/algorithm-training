package main

func maxProfit(prices []int) int {
    l := len(prices)
    min := 10000
    maxProfit := 0
    for i:=0;i<l;i++ {
        if prices[i] < min {
            min = prices[i]
            continue
        }
        if prices[i] - min > maxProfit {
            maxProfit = prices[i] - min
        }
    }
    return maxProfit
}

func main(){
	println(maxProfit([]int{7,1,5,3,6,4}))
	println(maxProfit([]int{7,6,4,3,1}))
}