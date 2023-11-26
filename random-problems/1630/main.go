package main

import (
	"fmt"
	"sort"
)

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	ans := make([]bool, len(l))
	for i := 0; i < len(l); i++ {
		if r[i]-l[i] <= 0 {
			ans[i] = false
			continue
		}
		ans[i] = true
		if r[i]-l[i] > 0 {
			aux := []int{}
			aux = append(aux, nums[l[i]:r[i]+1]...)
			sort.Ints(aux)
			diff := aux[1] - aux[0]
			for j := 1; j < len(aux); j++ {
				if aux[j]-aux[j-1] != diff {
					ans[i] = false
					break
				}
			}
		}
	}
	return ans
}

func main() {
	nums := []int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10}
	l := []int{0, 1, 6, 4, 8, 7}
	r := []int{4, 4, 9, 7, 9, 10}
	fmt.Println(checkArithmeticSubarrays(nums, l, r))
}
