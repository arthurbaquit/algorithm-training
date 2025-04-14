package main

func countGoodTriplets(arr []int, a int, b int, c int) int {
	n := len(arr)
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					count++
				}
			}
		}
	}
	return count
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func main() {
	println(countGoodTriplets([]int{3, 0, 1, 1, 9, 7}, 7, 2, 3)) // 4
	println(countGoodTriplets([]int{1, 1, 2, 2, 3}, 0, 0, 1))    // 0
}
