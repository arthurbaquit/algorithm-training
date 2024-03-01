package main

func canWinNim(n int) bool {
	return n%4 != 0
}

func main() {
	println(canWinNim(1))
	println(canWinNim(3))
	println(canWinNim(4))
	println(canWinNim(8))
}
