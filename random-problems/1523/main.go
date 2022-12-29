package main

func countOdds(low int, high int) int {
    low+=(low+1)%2
    if low > high {
        return 0
    }
    return (high - low) / 2 + 1
}

func main() {
	println(countOdds(3, 7))
	println(countOdds(8, 10))
	println(countOdds(1, 10))
	println(countOdds(1, 1))
	println(countOdds(1, 2))
	println(countOdds(2, 2))
}