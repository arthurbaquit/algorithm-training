package main

func lengthAfterTransformations(s string, t int) int {
	freq := make([]int, 26)
	mod := 1000000007
	for _, ch := range s {
		freq[ch-'a']++
	}
	for range t {
		aux := make([]int, 26)
		aux[0] = freq[25]
		aux[1] = (freq[25] + freq[0]) % mod
		for i := 2; i < 26; i++ {
			aux[i] = freq[i-1]
		}
		freq = aux
	}
	ans := 0
	for i := range 26 {
		ans = (ans + freq[i]) % mod
	}
	return ans
}

func main() {
	println(lengthAfterTransformations("abcyy", 2)) // 7
	println(lengthAfterTransformations("azbk", 1))  // 5
}
