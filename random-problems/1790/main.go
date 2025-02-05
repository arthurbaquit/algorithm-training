package main

func areAlmostEqual(s1 string, s2 string) bool {
	l := len(s1)
	first, second, count := -1, -1, 0
	for i := 0; i < l; i++ {
		if s1[i] != s2[i] {
			count++
			if count > 2 {
				return false
			}
			if first == -1 {
				first = i
				continue
			}
			second = i
		}
	}
	if count == 0 {
		return true
	}
	if first != -1 && second != -1 && s1[first] == s2[second] && s1[second] == s2[first] {
		return true
	}

	return false
}

func areAlmostEqualV2(s1 string, s2 string) bool {
	var diffS1, diffS2 byte
	alreadyChanged := false
	for i := range s1 {
		if s1[i] != s2[i] {
			if alreadyChanged {
				return false
			}
			if diffS1 == byte(0) {
				diffS1 = s1[i]
				diffS2 = s2[i]
				continue
			}
			if diffS1 == s2[i] && diffS2 == s1[i] {
				alreadyChanged = true
				continue
			}
			return false
		}
	}
	return !(diffS1 != byte(0) && !alreadyChanged)
}

func main() {
	println(areAlmostEqual("bank", "kanb"))       // true
	println(areAlmostEqual("attack", "defend"))   // false
	println(areAlmostEqual("kelb", "kelb"))       // true
	println(areAlmostEqual("aa", "ac"))           // false
	println(areAlmostEqualV2("bank", "kanb"))     // true
	println(areAlmostEqualV2("attack", "defend")) // false
	println(areAlmostEqualV2("kelb", "kelb"))     // true
	println(areAlmostEqualV2("aa", "ac"))         // false
}
