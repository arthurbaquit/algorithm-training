package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapStoT := make(map[byte]byte)
	mapLetter := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if _, ok := mapStoT[s[i]]; !ok {
			if _, ok := mapLetter[t[i]]; ok {
				return false
			}
			mapLetter[t[i]] = true
			mapStoT[s[i]] = t[i]

			continue
		}
		if mapStoT[s[i]] != t[i] {
			return false
		}
	}
	return true
}

func main() {
	println(isIsomorphic("egg", "add"))
	println(isIsomorphic("foo", "bar"))
	println(isIsomorphic("paper", "title"))
}
