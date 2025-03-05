package main

func coloredCells(n int) int64 {
	m := int64(n)
	return 1 + m*(m-1)*2
}

func main() {
	println(coloredCells(1)) // 1
	println(coloredCells(2)) // 5
	println(coloredCells(3)) // 13
}
