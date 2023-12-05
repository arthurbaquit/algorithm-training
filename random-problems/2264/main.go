package main

func largestGoodInteger(num string) string {
	i, m, e := 0, 1, 2
	ans := ""
	for e < len(num) {
		if num[i] == num[m] && num[m] == num[e] {
			if ans == "" {
				ans = num[i : e+1]
				continue
			}
			if num[i] > ans[0] {
				ans = num[i : e+1]
			}
		}
		i++
		m++
		e++
	}
	return ans
}

func main() {
	println(largestGoodInteger("1234567890"))
	println(largestGoodInteger("222312415"))
	println(largestGoodInteger("222317772415"))
}
