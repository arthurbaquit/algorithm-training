package main

import "fmt"

func myAtoi(s string) int {
	num := 0
	MaxInt, MinInt := (1<<31)-1, -1<<31
	neg, overflow, start := false, false, false
	for _, c := range s {
		if c == ' ' && !start {
			continue
		}
		if c == '+' && !start {
			start = true
			continue
		}
		if c == '0' {
			start = true
			if 10*num+0 <= MaxInt {
				num = 10*num + 0
				continue
			}
			overflow = true
			break
		}
		if c == '1' {
			start = true
			if 10*num+1 <= MaxInt {
				num = 10*num + 1
				continue
			}
			overflow = true
			break

		}
		if c == '2' {
			start = true
			if 10*num+2 <= MaxInt {
				num = 10*num + 2
				continue
			}
			overflow = true
			break

		}
		if c == '3' {
			start = true
			if 10*num+3 <= MaxInt {
				num = 10*num + 3
				continue
			}
			overflow = true
			break
		}
		if c == '4' {
			start = true
			if 10*num+4 <= MaxInt {
				num = 10*num + 4
				continue
			}
			overflow = true
			break
		}
		if c == '5' {
			start = true
			if 10*num+5 <= MaxInt {
				num = 10*num + 5
				continue
			}
			overflow = true
			break
		}
		if c == '6' {
			start = true
			if 10*num+6 <= MaxInt {
				num = 10*num + 6
				continue
			}
			overflow = true
			break
		}
		if c == '7' {
			start = true
			if 10*num+7 <= MaxInt {
				num = 10*num + 7
				continue
			}
			overflow = true
			break
		}
		if c == '8' {
			start = true
			if 10*num+8 <= MaxInt {
				num = 10*num + 8
				continue
			}
			overflow = true
			break
		}
		if c == '9' {
			start = true
			if 10*num+9 < MaxInt {
				num = 10*num + 9
				continue
			}
			overflow = true
			break
		}
		if c == '-' && !start {
			start = true
			neg = true
			continue
		}
		break
	}
	if neg {
		if overflow {
			return MinInt
		}
		return -1 * num
	}
	if overflow {
		return MaxInt
	}
	return num
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("3.14159"))
}
