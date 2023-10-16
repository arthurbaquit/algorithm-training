package main

import "fmt"

func getRow(rowIndex int) []int {
	res := []int{1}
	for i := 1; i <= rowIndex; i++ {
		res = append(res, res[i-1]*(rowIndex-i+1)/i)
	}
	return res
}

func main() {
	fmt.Println(getRow(3))
}
