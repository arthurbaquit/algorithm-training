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

func closestPrimesWithoutExtraSpace(left int, right int) []int {
	lastPrime := -1
	minDiff := 10000000
	cA, cB := -1, -1
	for i := left; i <= right; i++ {
		if isPrime(i) {
			if lastPrime != -1 {
				diff := i - lastPrime
				if diff < minDiff {
					minDiff = diff
					cA = lastPrime
					cB = i
				}
				if diff == 1 || diff == 2 {
					return []int{lastPrime, i}
				}
			}
			lastPrime = i
		}
	}
	if cA == -1 {
		return []int{-1, -1}
	}
	return []int{cA, cB}
}
func isPrime(x int) bool {
	if x == 1 {
		return false
	}
	if x == 2 || x == 3 {
		return true
	}
	if x%2 == 0 {
		return false
	}
	for i := 3; i*i <= x; i += 2 {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(closestPrimes(10, 19))                  // [11, 13]
	fmt.Println(closestPrimesWithoutExtraSpace(10, 19)) // [11, 13]
	fmt.Println(closestPrimesWithoutExtraSpace(4, 6))   // [-1, -1]
	fmt.Println(closestPrimes(4, 6))                    // [-1, -1]
}
