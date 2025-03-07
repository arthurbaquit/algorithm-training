package main

import "fmt"

func closestPrimes(left int, right int) []int {
	ans := []int{-1, -1}
	lastPrime := -1
	allPrimes := primes(right)
	for i := left; i <= right; i++ {
		if allPrimes[i] {
			if ans[0] == -1 {
				ans[0] = i
				lastPrime = i
				continue
			}
			if ans[1] == -1 {
				ans[1] = i
				lastPrime = i
				continue
			}
			if i-ans[1] < ans[1]-ans[0] {
				ans[0], ans[1] = ans[1], i
				lastPrime = i
				continue
			}
			if i-lastPrime < ans[1]-ans[0] {
				ans[0], ans[1] = lastPrime, i
			}
			lastPrime = i
		}
	}
	if ans[0] == -1 || ans[1] == -1 {
		return []int{-1, -1}
	}
	return ans
}
func primes(x int) []bool {
	allPrimes := make([]bool, x+1)
	for i := range allPrimes {
		allPrimes[i] = true
	}
	allPrimes[0] = false
	allPrimes[1] = false
	for i := 2; i*i <= x; i++ {
		if allPrimes[i] {
			for j := i * i; j <= x; j += i {
				allPrimes[j] = false
			}
		}
	}
	return allPrimes
}

func main() {
	fmt.Println(closestPrimes(10, 19)) // [11, 13]
	fmt.Println(closestPrimes(4, 6))   // [-1, -1]

}
