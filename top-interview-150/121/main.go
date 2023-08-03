package main

func maxProfit(prices []int) int {
	l := len(prices)
	min := prices[0]
	maxProfit := 0
	for i := 1; i < l; i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}
		if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}
	}
	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	println(maxProfit(prices))
}
