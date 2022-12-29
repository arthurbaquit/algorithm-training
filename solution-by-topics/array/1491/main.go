package main

func average(salary []int) float64 {
    min, max := int(^uint(0) >> 1), 0
    avg := 0
    for _, val := range(salary) {
        if val < min {
            min = val
        }
        if val > max {
            max = val
        }
        avg += val
    }
    return float64(avg - min - max)/ float64((len(salary)-2))
}

func main() {
	println(average([]int{4000,3000,1000,2000}))
	println(average([]int{1000,2000,3000}))
	println(average([]int{6000,5000,4000,3000,2000,1000}))
	println(average([]int{8000,9000,2000,3000,6000,1000}))
	println(average([]int{48000,59000,99000,13000,78000,45000,31000,17000,39000,37000,93000,77000,33000,28000,4000,54000,67000,6000,1000,11000}))
}