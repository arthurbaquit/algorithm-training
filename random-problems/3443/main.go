package main

func maxDistance(s string, k int) int {
	ans := 0
	north, south, east, west := 0, 0, 0, 0
	count := func(drt1, drt2, times int) int {
		return abs(drt1-drt2) + times*2
	}

	for _, dir := range s {
		switch dir {
		case 'N':
			north++
		case 'S':
			south++
		case 'E':
			east++
		case 'W':
			west++
		}
		times1 := min(min(north, south), k)      // modification times for N and S
		times2 := min(min(east, west), k-times1) // modification times for E and W
		current := count(north, south, times1) + count(east, west, times2)
		if current > ans {
			ans = current
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	println(maxDistance("NWSE", 1))   // 3
	println(maxDistance("NSWWEW", 3)) // 6
}
