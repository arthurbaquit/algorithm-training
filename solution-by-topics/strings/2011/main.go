package main

func finalValueAfterOperations(operations []string) int {
	ans := 0
	for _, v := range operations {
		if Contains(v, '+') {
			ans++
			continue
		}
		ans--
	}
	return ans
}

func Contains(value string, find rune) bool {
	for _, v := range value {
		if v == find {
			return true
		}
	}
	return false
}

func finalValueAfterOperationsWithLittleTrick(operations []string) int {
	ans := 0
	for _, v := range operations {
		if v[1] == '+' {
			ans++
			continue
		}
		ans--
	}
	return ans
}

func main() {
	println(finalValueAfterOperations([]string{"--X", "X++", "X++"}))
	println(finalValueAfterOperations([]string{"++X", "++X", "X++"}))
	println(finalValueAfterOperations([]string{"X++", "++X", "--X", "X--"}))
	println(finalValueAfterOperationsWithLittleTrick([]string{"--X", "X++", "X++"}))
	println(finalValueAfterOperationsWithLittleTrick([]string{"++X", "++X", "X++"}))
	println(finalValueAfterOperationsWithLittleTrick([]string{"X++", "++X", "--X", "X--"}))
}
