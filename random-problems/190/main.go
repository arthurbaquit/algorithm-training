package main

func reverseBits(num uint32) uint32 {
    var res uint32
    for i :=0; i<32; i++{
                res <<=1 
        res += num & 0x1
        num >>=1

    }
    return res
}

func reverseBitsWithoutForLoop(num uint32) uint32 {
	num = (num >> 16) | (num << 16)
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	return num
}

func main() {
	println(reverseBits(0b00000010100101000001111010011100))
	println(reverseBitsWithoutForLoop(0b11111111111111111111111111111101))
	println(reverseBits(0b00000010100101000001111010011100))
	println(reverseBitsWithoutForLoop(0b11111111111111111111111111111101))
}