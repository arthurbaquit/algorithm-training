package main

func minimumRecolors(blocks string, k int) int {
	wCount := 0
	for i := range k {
		if blocks[i] == 'W' {
			wCount++
		}
	}
	ans := wCount
	for r := k; r < len(blocks); r++ {
		if ans == 0 {
			return ans
		}
		if blocks[r-k] == 'W' {
			wCount--
		}
		if blocks[r] == 'W' {
			wCount++
		}
		ans = min(ans, wCount)
	}
	return ans
}

func main() {
	println(minimumRecolors("WBBWWBBWBW", 7)) // 3
	println(minimumRecolors("WBWBBBW", 2))    // 0
}
