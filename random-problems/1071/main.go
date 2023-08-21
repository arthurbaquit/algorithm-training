package main

func gcdOfStrings(str1 string, str2 string) string {
	// se existe um mínimo divisor comum, então tanto faz str1str2 ou str2str1
	if str1+str2 != str2+str1 {
		return ""
	}
	maxLen := GCD(len(str1), len(str2))
	return str1[:maxLen]
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {
	println(gcdOfStrings("ABCABC", "ABC"))
	println(gcdOfStrings("ABABAB", "ABAB"))
	println(gcdOfStrings("LEET", "CODE"))
}
