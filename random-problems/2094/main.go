package main

import "fmt"

func findEvenNumbers(digits []int) []int {
	var counts [10]int
	for _, digit := range digits {
		counts[digit]++
	}
	var ans []int
	for num := 100; num <= 998; num = num + 2 {
		numC := make([]int, 10)
		fstDig, secDig, trdDig := num%10, (num/10)%10, num/100
		numC[fstDig]++
		numC[secDig]++
		numC[trdDig]++
		if counts[fstDig] >= numC[fstDig] && counts[secDig] >= numC[secDig] && counts[trdDig] >= numC[trdDig] {
			ans = append(ans, num)
		}
	}
	return ans
}

func main() {
	fmt.Println(findEvenNumbers([]int{2, 1, 3, 0}))    // [102,120,130,132,210,230,302,310,312,320]
	fmt.Println(findEvenNumbers([]int{2, 2, 8, 8, 2})) // [222,228,282,288,822,828,882]
	fmt.Println(findEvenNumbers([]int{3, 7, 5}))       // []
}
