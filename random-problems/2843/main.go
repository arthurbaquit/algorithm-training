package main

func countSymmetricIntegers(low int, high int) int {
	count := 0
	for i := low; i <= high; i++ {
		if i < 10 {
			continue
		}
		if i < 100 {
			if i/10 == i%10 {
				count++
			}
		}
		if i < 1000 {
			continue
		}
		if i == 10000 {
			continue
		}
		if i/1000+(i/100)%10 == i%10+(i/10)%10 {
			count++
		}
	}
	return count
}

func main() {
	println(countSymmetricIntegers(1, 100))     // 9
	println(countSymmetricIntegers(1200, 1230)) // 4
}
