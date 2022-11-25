package main

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	i, j := 0, 0
	ii, jj := 0, 0
	l1, l2 := len(word1), len(word2)
	ll1, ll2 := len(word1[i]), len(word2[j])
	flag1, flag2 := false, false
	for {
		if word1[i][ii] != word2[j][jj] {
			return false
		}
		ii++
		jj++
		if ii > ll1-1 {
			i++
			ii = 0
			flag1 = true
		}
		if jj > ll2-1 {
			j++
			jj = 0
			flag2 = true
		}
		if i > l1-1 || j > l2-1 {
			break
		}
		if flag1 {
			ll1 = len(word1[i])
		}
		if flag2 {
			ll2 = len(word2[j])
		}
		flag1, flag2 = false, false

	}
	if i <= l1-1 || j <= l2-1 {
		return false
	}
	return true
}

func main() {
	word1 := []string{"ab", "c"}
	word2 := []string{"a", "bc"}
	println(arrayStringsAreEqual(word1, word2))
	word1 = []string{"a", "cb"}
	word2 = []string{"ab", "c"}
	println(arrayStringsAreEqual(word1, word2))
}
