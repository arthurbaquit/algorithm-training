package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	res := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				res = append(res, "FizzBuzz")
				continue
			}
			res = append(res, "Fizz")
			continue
		}
		if i%5 == 0 {
			res = append(res, "Buzz")
			continue
		}
		res = append(res, strconv.Itoa(i))
	}
	return res
}

func main() {
	fmt.Println(fizzBuzz(15))
	fmt.Println(fizzBuzz(20))
	fmt.Println(fizzBuzz(16))
	fmt.Println(fizzBuzz(25))
}
