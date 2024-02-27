package main

func countTime(time string) int {
	var res int
	hours := time[:2]
	min := time[3:]
	if hours[0] == '?' {
		if hours[1] == '?' {
			res = 24
		} else if hours[1] < '4' {
			res = 3
		} else {
			res = 2
		}
	} else if hours[1] == '?' {
		if hours[0] == '0' || hours[0] == '1' {
			res = 10
		} else {
			res = 4
		}
	} else {
		res = 1
	}
	if min[0] == '?' {
		if min[1] == '?' {
			res *= 60
		} else {
			res *= 6
		}
	} else {
		if min[1] == '?' {
			res *= 10
		}
	}
	return res
}

func main() {
	println(countTime("2?:?0"))
	println(countTime("1?:??"))
	println(countTime("??:??"))
}
