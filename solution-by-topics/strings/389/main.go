package main

func findTheDifference(s string, t string) byte {
    sMap := make(map[byte]int)
    for i:=0; i<len(s); i++{
        v := s[i]
        if _, ok := sMap[v]; ok{
            sMap[v]++
            continue
        }
        sMap[v] = 1
    }
    for i:=0; i<len(t); i++{
        v:=t[i]
        if _, ok := sMap[v]; ok{
            sMap[v]--
            if sMap[v]>=0 {continue}
        }
        return v
    }
    return t[len(t)-1]
}

func findTheDifferenceWithByteSum(s string, t string) byte {
	var sSum, tSum byte
	for i:=0; i<len(s); i++{
		sSum += s[i]
		tSum += t[i]
	}
	tSum += t[len(t)-1]
	return tSum - sSum
}

func findTheDifferenceWithXOR(s string, t string) byte {
	var res byte
	for i:=0; i<len(s); i++{
		res ^= s[i]
		res ^= t[i]
	}
	res ^= t[len(t)-1]
	return res
}
func main(){
	println(findTheDifference("abcd", "abcde"))
	println(findTheDifference("a", "aa"))
	println(findTheDifference("aa", "aaa"))
	println(findTheDifferenceWithByteSum("abcd", "abcde"))
	println(findTheDifferenceWithByteSum("a", "aa"))
	println(findTheDifferenceWithByteSum("aa", "aaa"))
	println(findTheDifferenceWithXOR("abcd", "abcde"))
	println(findTheDifferenceWithXOR("a", "aa"))
	println(findTheDifferenceWithXOR("aa", "aaa"))
}