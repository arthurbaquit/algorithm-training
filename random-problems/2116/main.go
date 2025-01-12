package main

func canBeValid(s string, locked string) bool {
	if len(s)%2 == 1 {
		return false
	}
	openB := 0
	unlocked := 0
	for i := range s {
		if locked[i] == '0' {
			unlocked++
		} else if s[i] == '(' {
			openB++
		} else if s[i] == ')' {
			if openB > 0 {
				openB--
			} else if unlocked > 0 {
				unlocked--
			} else {
				return false
			}
		}
	}
	balance := 0
	for i := len(s) - 1; i > -1; i-- {
		if locked[i] == '0' {
			balance--
			unlocked--
		} else if s[i] == '(' {
			balance++
			openB--
		} else if s[i] == ')' {
			balance--
		}
		if balance > 0 {
			return false
		}
		if unlocked == 0 && openB == 0 {
			break
		}
	}
	return !(openB > 0)
}

func main() {
	println(canBeValid("))()))", "010100"))                                                                                                             // true
	println(canBeValid("()()", "0000"))                                                                                                                 // true
	println(canBeValid("())(()(()(())()())(())((())(()())((())))))(((((((())(()))))(", "100011110110011011010111100111011101111110000101001101001111")) // false
	println(canBeValid("())()))()(()(((())(()()))))((((()()(()))", "1011101100010001001011000000110010100101"))                                         // true
	println(canBeValid(")(", "00"))                                                                                                                     // true
}
