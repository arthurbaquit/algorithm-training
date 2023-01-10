package main

func isPowerOfTwoWithBitsOps(n int) bool {
    if n <= 0 {return false}
    return n & (-n) == n
}

func isPowerOfTwo(n int) bool {
    var flag bool
    if n <= 0 {return false}
    for n > 0 {
        if n & 0x1 == 1 { if flag {return false}; flag = true}
        n = n>>1
    }
    return true
}

func main(){
	println(isPowerOfTwo(1))
	println(isPowerOfTwo(16))
	println(isPowerOfTwo(218))
	println(isPowerOfTwoWithBitsOps(1))
	println(isPowerOfTwoWithBitsOps(16))
	println(isPowerOfTwoWithBitsOps(218))
}