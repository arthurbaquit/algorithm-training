package main

func hammingWeight(num uint32) int {
    count := 0
    mask := uint32(1)
    for i:=0;i<32;i++{
        if (num&mask != 0) {
            count++
        }
        mask <<=1
    }
    return count
}

func OptimizedHammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num &= (num - 1)
	}
	return count
}

func main() {
	println("Hamming Weight")
	println(hammingWeight(0b00000000000000000000000000001011))
	println(hammingWeight(0b00000000000000000000000010000000))
	println(hammingWeight(0b11111111111111111111111111111101))
	println(hammingWeight(0b00000000000000000000000000000000))
	println(hammingWeight(0b11111111111111111111111111111111))
	println("Optimized Hamming Weight")
	println(OptimizedHammingWeight(0b00000000000000000000000000001011))
	println(OptimizedHammingWeight(0b00000000000000000000000010000000))
	println(OptimizedHammingWeight(0b11111111111111111111111111111101))
	println(OptimizedHammingWeight(0b00000000000000000000000000000000))
	println(OptimizedHammingWeight(0b11111111111111111111111111111111))
}